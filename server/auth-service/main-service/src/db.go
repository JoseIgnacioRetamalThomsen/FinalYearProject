// main package
package main

// provide acces to mysql database, use round_robin for load balancing read
// from slaves, all write are done to one master.

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

// database connection
var dbConn clientDB
var dbConnLB clientDBLoadBalancing

// Normal client context.
type clientDB struct {
	context *dbClientContext
}

// load balancer client context
type clientDBLoadBalancing struct {
	context *dbClientContextLoadBalancing
}

// Normal conection struc.
type dbClientContext struct {
	dbClient pb.UserAuthDBClient
	timeout  time.Duration
}

// load balancig conectio strucn
type dbClientContextLoadBalancing struct {
	dbClient pb.UserAuthDBClient
	timeout  time.Duration
}

// create normal connection
func newDBContext(endpoint string) (*dbClientContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &dbClientContext{
		dbClient: pb.NewUserAuthDBClient(userConn),
		timeout:  time.Second*MAX_CON_TIME,
	}
	return ctx, nil
}

// create load balancing conection
func newDBContextLoadBalancing() (*dbClientContextLoadBalancing, error) {
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", dbConnectionScheme, exampleServiceName),
		grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	ctx := &dbClientContextLoadBalancing{
		dbClient: pb.NewUserAuthDBClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

// Get users, user load balancing connection.
func getUser(email string) (string, []byte, []byte, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_CON_TIME)
	defer cancel()
	r, err := dbConnLB.context.dbClient.GetUser(ctx, &pb.UserDBRequest{Email: email})
	if err != nil {

		return "", nil, nil, errors.New("could not get user")
	}

	return r.GetEmail(), r.GetPasswordHash(), r.GetPasswordSalt(), nil
}

// add a new user , normal conection, always to master.
func addUser(email string, hashedPassword []byte, salt []byte) (string, int64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_CON_TIME)
	defer cancel()
	r, err := dbConn.context.dbClient.AddUser(ctx, &pb.UserDBRequest{Email: email, PasswordHash: hashedPassword, PasswordSalt: salt})
	if err != nil {
		return email, -1, err
	}
	return r.GetEmail(), r.GetId(), nil

}

// update user , normal connection, always to master.
func updateUser(email string, hash []byte, salt []byte) (string, []byte, []byte, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_CON_TIME)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdateUser(ctx, &pb.UserDBRequest{Email: email, PasswordHash: hash, PasswordSalt: salt})
	if err != nil {
		return "", nil, nil, err
	}
	return r.Email, r.PasswordHash, r.PasswordSalt, nil
}

// Create user seasion, normal conection, always to master.
func CreateSession(email string, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_CON_TIME)
	defer cancel()
	_, err := dbConn.context.dbClient.CreateSeassion(ctx, &pb.UserSessionRequest{Email: email, Token: token})
	if err != nil {
		return false, err
	}

	return true, nil

}

// Check token, use load balancing connection.
func CheckToken(email string, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_CON_TIME)
	defer cancel()
	res, err := dbConnLB.context.dbClient.GetSeassion(ctx, &pb.UserSessionRequest{Email: email, Token: token})
	if err != nil {
		return false, err
	}
	if res.Token == token {
		return true, nil
	}
	return false, nil
}

// Delete tocken, use normal connection, always to master,
func DeleteToken(email string, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*MAX_CON_TIME)
	defer cancel()
	res, err := dbConn.context.dbClient.DeleteSession(ctx, &pb.UserSessionRequest{Email: email, Token: token})
	if err != nil {
		return false, err
	}
	return res.Success, nil
}

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
			exampleServiceName: configuration.Dbs,
		},
	}
	r.start()
	return r, nil
}

// resolver implementation
func (*databasesResolverBuilder) Scheme() string { return dbConnectionScheme }

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
