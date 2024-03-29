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
        "github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/libvirt/internal/cobra/logs"
	"time"
	"os"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "gitlab.iblog.pro/cobra/libvirt/internal/cobra/protos"
        "gitlab.iblog.pro/cobra/libvirt/internal/cobra/system"
        grpcMetadata "google.golang.org/grpc/metadata"
        "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"

)

const (
	defaultName = "not-yet-set"
	defaultUUID = "2345234-5234523-45"
)

var (
	addr = flag.String("addr", "localhost:18080", "the address to connect to")
	name = flag.String("name", defaultName, "Not Yet Set")
        uuid = flag.String("uuid", defaultUUID, "UUID to greet")
)

var Version string

func main() {
	flag.Parse()

        system.VersionBuild(Version)

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
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in main()")
                os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewVirshClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

        md := grpcMetadata.Pairs("authorization", "bearer yolo")
	r, err := c.MachineHardReboot(metadata.MD(md).ToOutgoing(ctx), &pb.VirshRequest{Vmname: *name})

	if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in main()")
                os.Exit(1)
	}
	logs.Log.WithFields(logrus.Fields{ "message": fmt.Sprintf("Greeting: '%s' '%d' '%#v'", r.GetMessage(), r.GetCode(), r), }).Info("system message")

}
