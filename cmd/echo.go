package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/knabben/grpc-ex/damage"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Example echo GRPC service cli client",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []grpc.DialOption

		creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
		opts = append(opts, grpc.WithTransportCredentials(creds))
		conn, err := grpc.Dial(demoAddr, opts...)

		if err != nil {
			fmt.Println(err)
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := pb.NewDamageServiceClient(conn)

		msg, err := client.Damage(context.Background(),
			&pb.DamageMessage{Value: "data data data"})

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("data: ", msg)
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
