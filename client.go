package main

import (
	"os"

	"context"

	protos "github.com/harshthakur9030/grpc-primer/unary/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {

	log := logrus.New()
	log.Out = os.Stdout

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	// Create  a client for the service welcome
	wc := protos.NewWelcomeClient(conn)

	res, err := wc.World(context.Background(), &protos.Null{})
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	log.Info(res)

	// log.Info(res.GetName)

	// log.Info(res.GetNames)

	lc := protos.NewLookupClient(conn)

	// person := &protos.Person{Name: "CNCF"}
	// res, err = lc.Find(context.Background(), person)
	// if err != nil {
	// 	return err
	// }

	person := &protos.Person{Name: "grpc"}
	res, err = lc.Find(context.Background(), person)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			md := s.Details()[0]
			log.Infof("Metadata is %v", md)
			log.Error(err)
		}

	}

	log.Info(res.GetExist())

}
