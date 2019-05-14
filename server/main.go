package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	api "github.com/surenraju/grpc_rest_helloworld/greetingservice"
)

var (
	grpcPort = flag.Int("grpc-port", 8080, "The server grpc-port")
	httpPort = flag.Int("http-port", 8080, "The server http-port")
)

type greetServiceServer struct {
}

func main() {

	flag.Parse()
	ctx := context.Background()
	go func() {
		_ = RunRestServer(ctx, strconv.Itoa(*grpcPort), strconv.Itoa(*httpPort))
	}()

	_ = RunGrpcServer(ctx, strconv.Itoa(*grpcPort))
}

func (s *greetServiceServer) Greet(ctx context.Context, req *api.GreetRequest) (*api.GreetResponse, error) {
	return &api.GreetResponse{Greeting: fmt.Sprintf("Hello %s", req.Name)}, nil
}

//Run gRPC server
func RunGrpcServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	api.RegisterGreetServiceServer(server, &greetServiceServer{})

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
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}

// run REST server
func RunRestServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterGreetServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
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

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
