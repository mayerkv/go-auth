package main

import (
	"context"
	"fmt"
	"log"

	grpc_service2 "github.com/mayerkv/go-auth/internal/grpc-service"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpc_service2.NewAuthServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &grpc_service2.CreateAccountRequest{
		Email:    "foo1@bar.com",
		Password: "test",
		UserId:   "2",
		Role:     grpc_service2.AccountRole_USER,
	}
	account, err := client.CreateAccount(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.String())
}
