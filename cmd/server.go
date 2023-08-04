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
	"net"
	"os"
        "github.com/sirupsen/logrus"
        logs "iblog.pro/cobra/logs"
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
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in main()")
		os.Exit(1)
	}
	s := grpc.NewServer()
	pb.RegisterVirshServer(s, &Virsh.Server{})
	logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("server listening at %v", lis.Addr()), }).Error("System message in main()")
	if err := s.Serve(lis); err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in main()")
	}


}
