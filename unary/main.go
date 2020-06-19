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

	protos.RegisterLookupServer(gs, &lookup{})

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

// // First Example
// func (w *welcome) World(ctx context.Context, null *protos.Null) (*protos.Person, error) {

// 	name := "CNCF and Hashicorp User group"

// 	return &protos.Person{
// 		Name: name,
// 	}, nil
// }

// Second example
func (w *welcome) World(ctx context.Context, null *protos.Null) (*protos.Person, error) {

	name := "CNCF and Hashicorp User group"

	names := make([]string, 0)

	names = append(names, "CNCF")

	names = append(names, "Hashicorp")

	return &protos.Person{
		Name:  name,
		Names: names,
	}, nil
}

// Third example

type lookup struct {
}

func (l *lookup) Find(ctx context.Context, p *protos.Person) (*protos.Person, error) {
	names := make([]string, 0)
	var exist bool

	names = append(names, "Hashicorp")
	names = append(names, "CNCF")

	for _, n := range names {
		if n == p.Name {
			exist = true
		}
	}

	return &protos.Person{
		Exist: exist,
	}, nil
}
