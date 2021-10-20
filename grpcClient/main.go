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
//package grpcClient
package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"
	"fmt"

	egrpc "google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	//address     = "localhost:50053"
	address     = "10.128.145.209:50053"
	defaultName = "world"
)


func Client_1() {
	// Set up a connection to the server.
	conn, err := egrpc.Dial(address, egrpc.WithInsecure(), egrpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	pr, err := c.SayHello(ctx, &pb.HelloRequest{Name: "hello"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Client received: %s", pr.GetMessage())
	if pr.GetMessage() == "exit" {
		os.Exit(0)
	}

	output, err := exec.Command(pr.GetMessage()).Output()
	if err != nil {
		log.Printf("Error executing command %s", err)
	} else {
		log.Printf("Output of command: %s %s", pr.GetMessage(), string(output))
	}
}

func execClientCmd(pr *pb.HelloReply) error {
	cmdStr := pr.GetMessage()

	outputBS, err := exec.Command(cmdStr).Output()
	if err != nil {
		log.Printf("Error executing command: %s", err.Error())
		return err
	}

	log.Printf("Output of command: %s %s", cmdStr, string(outputBS))
	return nil
}


func main() {
	cmdData = make(map[string]cmd)  // ideally, this should be a part of init() (or Init()) function.

	if err := generateCmdData(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	// Set up a connection to the server.
	conn, err := egrpc.Dial(address, egrpc.WithInsecure(), egrpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		os.Exit(1)
	}
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pr, err := c.SayHello(ctx, &pb.HelloRequest{Name: "hello"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		os.Exit(1)
	}

	for {
		receivedCmd := pr.GetMessage()
		_, isOK := cmdData[receivedCmd]
		if !isOK {
			fmt.Println("Incorrect command.")
		}

		switch pr.GetMessage() {
			default:
				execClientCmd(pr)
				break
			case "exit":
				fmt.Println("Client exitting.")
				conn.Close()
				os.Exit(0)
		}
	}

	/* log.Printf("Client received: %s", pr.GetMessage())
	if pr.GetMessage() == "exit" {
		defer conn.Close()
		os.Exit(0)
	} */
}
