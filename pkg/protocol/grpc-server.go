package protocol

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"bandi.com/main/pkg/auth"
	"bandi.com/main/pkg/data"
	"bandi.com/main/pkg/service"
)

// RunGrpcServer runs the grpc server
func RunGrpcServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			auth.StreamServerInterceptor(auth.CheckForBearerToken),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			auth.UnaryServerInterceptor(auth.CheckForBearerToken),
		)),
	)
	data.RegisterUserServiceServer(server, service.NewUserService())
	data.RegisterBandiServiceServer(server, service.NewBandiUserService())

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("starting gRPC server at port =%v...\n", port)
	return server.Serve(listen)
}
