package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/endzyme/telephone/protobuf/krbrnrtr"
	"google.golang.org/grpc"
)

func main() {
	var webEndpoint string

	webEndpoint = os.Args[1]

	var grpcOpts []grpc.DialOption

	grpcOpts = append(grpcOpts, grpc.WithInsecure())
	conn, err := grpc.Dial(webEndpoint, grpcOpts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := krbrnrtr.NewKrbrnrtrClient(conn)

	req := krbrnrtr.CommandRequest{
		Action:    "curl",
		Arguments: []string{"-L", "https://google.com"},
	}

	resp, err := client.RunActionSyncResult(context.Background(), &req)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := resp.Recv()
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Println(string(msg.GetStatusMessage()))
	}

}
