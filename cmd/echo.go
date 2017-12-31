package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/knabben/grpc-ex/damage"
)

var (
	remoteHost string
	echoCmd    = &cobra.Command{
		Use:   "echo",
		Short: "Example echo GRPC service cli client",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Connecting on", remoteHost)
			conn, err := grpc.Dial(remoteHost, grpc.WithInsecure())
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
)

func init() {
	echoCmd.Flags().StringVar(&remoteHost, "grpc", "localhost:9090", "Remote Host")
	rootCmd.AddCommand(echoCmd)
}
