package main

import (
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

type profileServer struct {
	context *ProfileServiceContext
}

type ProfileServiceContext struct {
	dbClient pb.ProfilesClient
	timeout  time.Duration
}


var ProfSerConn profileServer

// create connection
func newProfilesServiceContext(endpoint string) (*ProfileServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &ProfileServiceContext{
		dbClient: pb.NewProfilesClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

// create load balancing conection
func newProfilesContextLoadBalancing() (*ProfileServiceContext, error) {
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", profilesConnectionScheme, profilesServiceName),
		grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	ctx := &ProfileServiceContext{
		dbClient: pb.NewProfilesClient(userConn),
		timeout:  time.Second*10,
	}
	return ctx, nil
}

// configuration
const (
	profilesConnectionScheme = "profilesConnectionScheme"
	profilesServiceName      = "github.com.joseignacioretamalthomsen.wcity.profiles"
)

//resolver
type profilesResolverBuilder struct{}

//resolver struct
type profilesResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

// resolver implementation
func (*profilesResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &profilesResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			profilesServiceName: configuration.Profiles,
		},
	}
	r.start()
	return r, nil
}

// resolver implementation
func (*profilesResolverBuilder) Scheme() string { return profilesConnectionScheme }
// resolver implementation
func (r *profilesResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

// resolver implementation
func (*profilesResolver) ResolveNow(o resolver.ResolveNowOptions) {}

// resolver implementation
func (*profilesResolver) Close() {}

//init resolvers.
func init() {
	resolver.Register(&profilesResolverBuilder{})
}


