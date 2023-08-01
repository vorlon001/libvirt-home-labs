/*
 *
 * Copyright 2023, vorlon
 *
 * GRPC server in Libvirt
 *
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "iblog.pro/condor/protos"
	Virsh "iblog.pro/cobra/core/virshgrpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)


func main() {

        ctx := context.Background()
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVirshServer(s, &Virsh.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}
