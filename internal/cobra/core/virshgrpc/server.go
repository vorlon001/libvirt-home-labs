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

	pb "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/protos"
        VirtualMachine "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/core/virtualmachine"

)



func NewServer() *Server {
	return &Server{}
}
type Server struct {
	pb.UnimplementedVirshServer
}


func (s *Server) MachineState(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
	domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
		return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineDelete(in.GetVmname())
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


        log.Printf("func MachineState(). Received: %v %v", in.GetXml())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineDelete(in.GetXml())
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



        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineDelete(in.GetVmname())
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


        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineSoftReboot(in.GetVmname())
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


        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineHardReboot(in.GetVmname())
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


        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineShutdown(in.GetVmname())
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


        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineStart(in.GetVmname())
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


        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachinePause(in.GetVmname())
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


        log.Printf("func MachineState(). Received: %v %v", in.GetVmname())
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                log.Printf("ERROR:%s\n", err)
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineResume(in.GetVmname())
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


