package main

import (
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

/*
RESOLVER IMPLEMENTATION
*/
var dbConnLB clientDBLoadBalancing

// load balancer client context
type clientDBLoadBalancing struct {
	context *dbClientContextLoadBalancing
}
type dbClientContextLoadBalancing struct {
	dbClient pb.UserAuthenticationClient
	timeout  time.Duration
}

// create load balancing conection
func newDBContextLoadBalancing() (*dbClientContextLoadBalancing, error) {
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", authConnectionScheme, authServiceName),
		grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	ctx := &dbClientContextLoadBalancing{
		dbClient: pb.NewUserAuthenticationClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

// configuration
const (
	authConnectionScheme = "authConnectionScheme"
	authServiceName      = "github.com.joseignacioretamalthomsen.wcity.auth"
)

//resolver
type databasesResolverBuilder struct{}

//resolver struct
type databaseResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

// resolver implementation
func (*databasesResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &databaseResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			authServiceName: configuration.Auth,
		},
	}
	r.start()
	return r, nil
}

// resolver implementation
func (*databasesResolverBuilder) Scheme() string { return authConnectionScheme }
// resolver implementation
func (r *databaseResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

// resolver implementation
func (*databaseResolver) ResolveNow(o resolver.ResolveNowOptions) {}

// resolver implementation
func (*databaseResolver) Close() {}

//init resolvers.
func init() {
	resolver.Register(&databasesResolverBuilder{})
}


