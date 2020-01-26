package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/joseignacioretamalthomsen/wcity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

type clientDB struct {
	context *dbClientContext
}

type clientDBLoadBalancing struct {
	context *dbClientContext
}

type dbClientContext struct {
	dbClient pb.UserAuthDBClient
	timeout time.Duration
}
type dbClientContextLoadBalancing struct {
	dbClient pb.UserAuthDBClient
	timeout time.Duration
}

// create connection
func newDBContext(endpoint string) (*dbClientContext, error) {
	userConn, err := grpc.Dial(
		endpoint,
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	ctx := &dbClientContext{
		dbClient: pb.NewUserAuthDBClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

func newDBContextLoadBalancing() (*dbClientContext, error) {
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", dbConnectionScheme, exampleServiceName),
		grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	ctx := &dbClientContext{
		dbClient: pb.NewUserAuthDBClient(userConn),
		timeout:  time.Second,
	}
	return ctx, nil
}

func getUser(email string) (string,[]byte,[]byte,error){

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConnLB.context.dbClient.GetUser(ctx, &pb.UserDBRequest{Email: email})
	if err != nil {

		return "",nil,nil, errors.New("could not get user")
	}

	return r.GetEmail(),r.GetPasswordHash(), r.GetPasswordSalt(),nil
}



func addUser(email string, hashedPassword []byte, salt []byte) (string,int64,error){

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.AddUser(ctx, &pb.UserDBRequest{Email: email, PasswordHash: hashedPassword,PasswordSalt:salt})
	if err != nil {
		return email,-1, err
	}
	return r.GetEmail(),r.GetId(),nil

}

func updateUser(email string, hash []byte, salt []byte) (string,[]byte,[]byte,error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := dbConn.context.dbClient.UpdateUser(ctx, &pb.UserDBRequest{Email: email,PasswordHash:hash,PasswordSalt:salt})
	if err != nil {
		return "", nil,nil,err
	}
	return r.Email,r.PasswordHash,r.PasswordSalt,nil
}

func CreateSession(email string, token string) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := dbConn.context.dbClient.CreateSeassion(ctx,&pb.UserSessionRequest{Email:email,Token:token})
	if err != nil{
       return false,err
	}

    return true,nil

}

func CheckToken(email string, token string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res,err := dbConn.context.dbClient.GetSeassion(ctx, &pb.UserSessionRequest{Email:email,Token:token})
	if err!=nil{
		return false,err
	}
	if res.Token == token{
		return true,nil
	}
	return false, nil
}

func DeleteToken(email string,token string) (bool,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res,err := dbConn.context.dbClient.DeleteSession(ctx, &pb.UserSessionRequest{Email:email,Token:token})
	if err !=nil{
		return false, err
	}
	return res.Success , nil
}

//resolver
type databasesResolverBuilder struct{}

func (*databasesResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
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
func (*databasesResolverBuilder) Scheme() string { return dbConnectionScheme }

type databaseResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *databaseResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (*databaseResolver) ResolveNow(o resolver.ResolveNowOption) {}
func (*databaseResolver) Close(){}
