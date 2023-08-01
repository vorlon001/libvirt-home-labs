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
        VirtualMachine "iblog.pro/cobra/core/virtualmachine"

)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedVirshServer
}


func (s *server) MachineState(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineState(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachineDelete(in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *server) MachineCreate(ctx context.Context, in *pb.VirshCreateRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineCreate(). Received: %v", in.GetXml())

        message, err := VirtualMachine.Virtinit().VirtualMachineDelete(in.GetXml())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *server) MachineDelete(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineSoftReboot(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachineDelete(in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *server) MachineSoftReboot(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineSoftReboot(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachineSoftReboot(in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *server) MachineHardReboot(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

	log.Printf("func MachineHardReboot(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachineHardReboot(in.GetVmname())
	messageSend := ""
	var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
		messageSend = fmt.Sprintf("%s", err)
		code = 500
        } else {
                log.Printf("DONE:%s\n",message)
		messageSend = fmt.Sprintf("%s", message)
		code = 200
        }


	return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *server) MachineShutdown(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineShutdown(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachineShutdown(in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *server) MachineStart(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineHardReboot(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachineStart(in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}



func (s *server) MachinePause(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachinePause(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachinePause(in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *server) MachineResume(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachinePause(). Received: %v", in.GetVmname())

        message, err := VirtualMachine.Virtinit().VirtualMachineResume(in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                log.Printf("DONE:%s\n",message)
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


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
	pb.RegisterVirshServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}
