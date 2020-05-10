package main

import (
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

type postService struct {
	context *postServiceContext
}

type postServiceContext struct {
	dbClient pb.PostsServiceClient
	timeout time.Duration
}
var serviceConn postService

// create connection
func newPostServiceContext(endpoint string) (*postServiceContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &postServiceContext{
		dbClient: pb.NewPostsServiceClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}
// create load balancing conection
func newPostContextLoadBalancing() (*postServiceContext, error) {
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", postConnectionScheme, postServiceName),
		grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	ctx := &postServiceContext{
		dbClient: pb.NewPostsServiceClient(userConn),
		timeout:  time.Second*10,
	}
	return ctx, nil
}


// configuration
const (
	postConnectionScheme = "postConnectionScheme"
	postServiceName      = "github.com.joseignacioretamalthomsen.wcity.post"
)

//resolver
type postResolverBuilder struct{}

//resolver struct
type postResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

// resolver implementation
func (*postResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &postResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			postServiceName: configuration.Post,
		},
	}
	r.start()
	return r, nil
}

// resolver implementation
func (*postResolverBuilder) Scheme() string { return postConnectionScheme }
// resolver implementation
func (r *postResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

// resolver implementation
func (*postResolver) ResolveNow(o resolver.ResolveNowOptions) {}

// resolver implementation
func (*postResolver) Close() {}

//init resolvers.
func init() {
	resolver.Register(&postResolverBuilder{})
}



