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
package main

import (
	"context"
	"flag"
	"log"
	"time"
//	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "iblog.pro/condor/protos"
//        VirtualMachine "iblog.pro/cobra/core/virtualmachine"

// google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "not-yet-set"
	defaultUUID = "2345234-5234523-45"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Not Yet Set")
        uuid = flag.String("uuid", defaultUUID, "UUID to greet")
)

func main() {
	flag.Parse()

/*	message, err := VirtualMachine.Virtinit().VirtualMachineHardReboot("node150")
        if err != nil {
		fmt.Printf("ERROR:%s\n", err)
	} else {
	        fmt.Printf("DONE:%s\n",message)
	}

	return*/

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVirshClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MachineHardReboot(ctx, &pb.VirshRequest{Vmname: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: '%s' '%d' '%#v'", r.GetMessage(), r.GetCode(), r)

}
