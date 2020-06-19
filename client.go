package main

import (
	"os"

	"context"

	protos "github.com/harshthakur9030/grpc-primer/unary/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	person := &protos.Person{Name: "CNCF"}
	res, err = lc.Find(context.Background(), person)

	log.Info(res.GetExist())

}
