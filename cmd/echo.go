package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/knabben/grpc/damage"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Example echo GRPC service cli client",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":9090", grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()

		client := pb.NewDamageServiceClient(conn)
		msg, err := client.Damage(context.Background(),
			&pb.DamageMessage{Value: "data data data"})
		fmt.Println("data: ", msg)
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
