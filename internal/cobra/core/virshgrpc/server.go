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
        "github.com/sirupsen/logrus"

        logs "gitlab.iblog.pro/cobra/libvirt/internal/cobra/logs"
	pb "gitlab.iblog.pro/cobra/libvirt/internal/cobra/protos"
        VirtualMachine "gitlab.iblog.pro/cobra/libvirt/internal/cobra/core/virtualmachine"

)



func NewServer() *Server {
	return &Server{}
}
type Server struct {
	pb.UnimplementedVirshServer
}



func (s *Server) MachineMigrate(ctx context.Context, in *pb.VirshMachineMigrate) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), "in.GetTomove()": fmt.Sprintf("%v", in.GetTomove()),}).Error("System Error in MachineState()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineState()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineMigrate(in.GetVmname(), in.GetTomove())
        if err != nil {
                logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineState()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
                logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineState()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil

}

func (s *Server) MachineState(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

	logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachineState()")
        messageSend := ""
        var code int32 = 0;
	domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineState()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
		return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineState(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineState()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineState()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *Server) MachineCreate(ctx context.Context, in *pb.VirshCreateRequest) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetXml()), }).Error("System Error in MachineCreate()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineCreate()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineCreate(in.GetXml())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineCreate()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineCreate()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *Server) MachineDelete(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {



        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachineDelete()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineDelete()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineDelete(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineDelete()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineDelete()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *Server) MachineSoftReboot(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachineSoftReboot()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineSoftReboot()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineSoftReboot(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineSoftReboot()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineSoftReboot()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *Server) MachineHardReboot(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachineHardReboot()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineHardReboot()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineHardReboot(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineHardReboot()")
		messageSend = fmt.Sprintf("%s", err)
		code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineHardReboot()")
		messageSend = fmt.Sprintf("%s", message)
		code = 200
        }


	return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *Server) MachineShutdown(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachineShutdown()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineShutdown()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineShutdown(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineShutdown()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineShutdown()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *Server) MachineStart(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachineStart()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineStart()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineStart(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineStart()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{ "message": message, }).Error("System Error in MachineStart()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}



func (s *Server) MachinePause(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachinePause()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachinePause()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachinePause(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachinePause()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{  "message": message, }).Error("System Error in MachinePause()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


func (s *Server) MachineResume(ctx context.Context, in *pb.VirshRequest) (*pb.VirshReply, error) {

        logs.Log.WithFields(logrus.Fields{ "in.GetVmname()": fmt.Sprintf("%v", in.GetVmname()), }).Error("System Error in MachineResume()")
        messageSend := ""
        var code int32 = 0;
        domain, err := VirtualMachine.Virtinit()
        if err != nil {
                logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineResume()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
                return &pb.VirshReply{Message: messageSend, Code: code}, nil
        }

        message, err := domain.VirtualMachineResume(in.GetVmname())
        if err != nil {
		logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in MachineResume()")
                messageSend = fmt.Sprintf("%s", err)
                code = 500
        } else {
		logs.Log.WithFields(logrus.Fields{  "message": message, }).Error("System Error in MachineResume()")
                messageSend = fmt.Sprintf("%s", message)
                code = 200
        }


        return &pb.VirshReply{Message: messageSend, Code: code}, nil
}


