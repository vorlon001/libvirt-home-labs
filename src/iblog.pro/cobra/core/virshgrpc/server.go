/*
 *
 * Copyright 2023, vorlon
 *
 * GRPC server in Libvirt
 *
 */

package server

import (
	"context"
	"fmt"
	"log"

	pb "iblog.pro/condor/protos"
        VirtualMachine "iblog.pro/cobra/core/virtualmachine"

)


type Server struct {
	pb.UnimplementedVirshServer
}


func (s *Server) MachineState(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())

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


func (s *Server) MachineCreate(ctx context.Context, in *pb.VirshCreateRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineCreate(). Received: %v %v", in.GetXml())

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


func (s *Server) MachineDelete(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineSoftReboot(). Received: %v %v", in.GetVmname())

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


func (s *Server) MachineSoftReboot(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineSoftReboot(). Received: %v %v", in.GetVmname())

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


func (s *Server) MachineHardReboot(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

	log.Printf("func MachineHardReboot(). Received: %v %v", in.GetVmname())

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


func (s *Server) MachineShutdown(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineShutdown(). Received: %v %v", in.GetVmname())

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


func (s *Server) MachineStart(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineHardReboot(). Received: %v %v", in.GetVmname())

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



func (s *Server) MachinePause(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachinePause(). Received: %v %v", in.GetVmname())

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


func (s *Server) MachineResume(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachinePause(). Received: %v %v", in.GetVmname())

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


