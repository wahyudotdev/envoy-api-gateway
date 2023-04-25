package main

import (
	"fmt"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"time"
	"try-gcp-grpc/services"
	"try-gcp-grpc/services/bookstore"
)

func main() {
	logrusEntry := logrus.NewEntry(logrus.New())
	logrusOpts := []grpclogrus.Option{
		grpclogrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_s", duration.Nanoseconds()
		}),
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(middleware.ChainUnaryServer(
		grpclogrus.UnaryServerInterceptor(logrusEntry, logrusOpts...),
	)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	service := services.NewHelloService()
	bookstore.RegisterBookStoreServiceServer(server, service)
	reflection.Register(server)
	log.Println("grpc server started")
	log.Fatal(server.Serve(l))
}
