package main

import (
	"net"
	"os"

	"context"

	protos "github.com/harshthakur9030/grpc-primer/unary/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log := logrus.New()
	log.Out = os.Stdout
	// Creating a grpc server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// This helps clients determine which services are available to call( like robots.txt but for gRPC)
	reflection.Register(gs)

	// Register the name of the service to the server
	protos.RegisterWelcomeServer(gs, &welcome{})

	l, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("Unable to listen %v", err)
	}

	// Listen for requests
	log.Info("Starting server at 9090")
	gs.Serve(l)
}

type welcome struct {
}

// First Example
func (w *welcome) World(ctx context.Context, null *protos.Null) (*protos.Person, error) {

	name := "CNCF Mentees meetup"

	return &protos.Person{
		Name: name,
	}, nil
}
