package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mayerkv/go-auth/domain"
	"github.com/mayerkv/go-auth/grpc-service"
	"github.com/mayerkv/go-auth/http-service"
	"github.com/mayerkv/go-auth/repository"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	authConfig := domain.NewAuthConfig(privateKey, jwt.SigningMethodPS256, "auth-service", 15*time.Minute, 24*time.Hour)

	accountRepository := repository.NewInMemoryAccountRepository()
	passwordEncoder := domain.NewBCryptPasswordEncoder()
	authService := domain.NewAuthService(accountRepository, passwordEncoder, *authConfig)
	accountService := domain.NewAccountService(passwordEncoder, accountRepository)

	account := domain.CreateAccount("foo@bar.com", "test", "1", domain.AccountRoleAdmin, passwordEncoder)
	accountRepository.Save(account)

	srv := grpc_service.NewAuthServiceServerImpl(accountService)

	g := newGroup()
	runHttpSever(g, authService)
	runGrpcServer(g, srv)
	runGraceful(g)

	if err = g.run(); err != nil {
		log.Fatal(err)
	}
}

func runGraceful(g *group) {
	cancelInterrupt := make(chan struct{})
	g.add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func runHttpSever(g *group, authService *domain.AuthService) {
	controller := http_service.NewAuthController(authService)
	router := http_service.CreateRouter(controller, authService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	g.add(func() error {
		return srv.ListenAndServe()
	}, func(err error) {
		log.Printf("http err: %s", err)
		srv.Shutdown(context.Background())
	})
}

func runGrpcServer(g *group, srv grpc_service.AuthServiceServer) {
	var grpcServer *grpc.Server
	var lis net.Listener

	g.add(func() error {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			return err
		}

		var opts []grpc.ServerOption
		grpcServer = grpc.NewServer(opts...)
		grpc_service.RegisterAuthServiceServer(grpcServer, srv)

		return grpcServer.Serve(lis)
	}, func(err error) {
		log.Printf("grpc err: %s", err)
		if grpcServer != nil {
			grpcServer.Stop()
		}
		if lis != nil {
			lis.Close()
		}
	})

}

type actor struct {
	execute   func() error
	interrupt func(error)
}

type group struct {
	actors []actor
}

func newGroup() *group {
	return &group{actors: []actor{}}
}

func (g *group) add(execute func() error, interrupt func(err error)) {
	g.actors = append(g.actors, actor{execute, interrupt})
}

func (g *group) run() error {
	if len(g.actors) == 0 {
		return nil
	}

	errors := make(chan error, len(g.actors))
	for _, a := range g.actors {
		go func(a actor) {
			errors <- a.execute()
		}(a)
	}

	err := <-errors

	for _, a := range g.actors {
		a.interrupt(err)
	}

	for i := 1; i < cap(errors); i++ {
		<-errors
	}

	return err
}
