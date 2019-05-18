package api

import (
	"fmt"
	"net"
	"os"

	"github.com/endzyme/slipway/pkg/cluster"
	"github.com/endzyme/slipway/protobuf/slipway"
	"google.golang.org/grpc"
)

//RunActionSyncResult asdf
// func (s Server) RunActionSyncResult(req *krbrnrtr.CommandRequest, resultServer krbrnrtr.Krbrnrtr_RunActionSyncResultServer) error {
// 	cmdOptions := cmd.Options{
// 		Buffered:  false,
// 		Streaming: true,
// 	}

// 	action := cmd.NewCmdOptions(cmdOptions, req.Action, req.Arguments...)
// 	receiver := "self"

// 	go func() {
// 		var msg krbrnrtr.CommandResponse
// 		for {
// 			select {
// 			case line := <-action.Stdout:
// 				msg.ReceiverNode = receiver
// 				msg.StatusMessage = []byte(fmt.Sprintf("%s: %s", time.Now().String(), line))
// 			case line := <-action.Stderr:
// 				msg.ReceiverNode = receiver
// 				msg.StatusMessage = []byte(line)
// 			}
// 			resultServer.Send(&msg)
// 		}
// 	}()

// 	<-action.Start()

// 	for len(action.Stdout) > 0 || len(action.Stderr) > 0 {
// 		time.Sleep(10 * time.Millisecond)
// 	}

// 	log.Printf("Finished command %v", action.Name)

// 	if action.Status().Exit > 0 {
// 		return fmt.Errorf("Error found executing thing. Exit Code: %d", action.Status().Exit)
// 	}

// 	return nil
// }

// ServeGRPC returns the bind port and http server to listen
func ServeGRPC(bindAddress string, slipwayCluster cluster.SlipwayCluster, gracefulStop chan os.Signal) error {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	listener, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return err
	}

	kubernetesClusterHandler := KubernetesService{}
	kubernetesNodeHandler := KubernetesNodeHandler{}
	clusterHandler := ClusterHandler{slipwayCluster: slipwayCluster}

	slipway.RegisterKubernetesClusterServer(grpcServer, kubernetesClusterHandler)
	slipway.RegisterKubernetesNodeServer(grpcServer, kubernetesNodeHandler)
	slipway.RegisterClusterServer(grpcServer, clusterHandler)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v\n", sig)
		fmt.Println("Stopping API Server")
		grpcServer.GracefulStop()
	}()

	return grpcServer.Serve(listener)
}
