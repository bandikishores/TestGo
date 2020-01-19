package protocol

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"bandi.com/main/pkg/data"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// RunRestServer runs HTTP/REST gateway
func RunRestServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := data.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Printf("starting skygraph HTTP/REST gateway at port=%v...\n", httpPort)
	return srv.ListenAndServe()
}
