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
	"os/signal"


        "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

        logs "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/logs"
	"google.golang.org/grpc"
	pb "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/protos"
	Virsh "gitlab.iblog.pro/cobra/libvirtgrpc/internal/cobra/core/virshgrpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)


func main() {

	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	eg, _ := errgroup.WithContext(ctx)


        srv := grpc.NewServer()


	eg.Go(func() error {

		logs.Log.WithFields(logrus.Fields{ "message": fmt.Sprintf("starting server 0.0.0.0:%d", *port), }).Info("system message")

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
                	logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("%v", err), }).Error("System Error in main()")
	                os.Exit(1)
        	}
	        pb.RegisterVirshServer(srv, Virsh.NewServer())
        	logs.Log.WithFields(logrus.Fields{ "error": fmt.Sprintf("server listening at %v", lis.Addr()), }).Info("System message in main()")

		return srv.Serve(lis)
	})



	eg.Go(func() error {
		// shutdown server
		<-ctx.Done()
                logs.Log.Info("stopping server")
                defer logs.Log.Info("server stopped")

		srv.Stop()
		return nil
	})

	_ = eg.Wait()

}
