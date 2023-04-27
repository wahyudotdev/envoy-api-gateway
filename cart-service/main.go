package main

import (
	"cart-service/services"
	"cart-service/services/cart"
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
	service := services.NewCartService()
	cart.RegisterCartServiceServer(server, service)
	reflection.Register(server)
	log.Println("cart service started")
	log.Fatal(server.Serve(l))
}
