package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
)

func newGRPCConnection(o Options) (*grpc.ClientConn, error) {
	r, _ := naming.NewDNSResolverWithFreq(time.Second * 1)
	balancer := grpc.RoundRobin(r)

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	conn, err := grpc.DialContext(ctx, o.Address, grpc.WithInsecure(), grpc.WithBalancer(balancer), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}

	return conn, err
}
