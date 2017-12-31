package cmd

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"

	pb "github.com/knabben/grpc-ex/damage"
	"github.com/spf13/cobra"
)

type server struct{}

func newServer() pb.DamageServiceServer {
	return new(server)
}

func (s *server) Damage(ctx context.Context, msg *pb.DamageMessage) (*pb.DamageMessage, error) {
	fmt.Println("Received: ", msg)
	return &pb.DamageMessage{Value: msg.Value}, nil
}

var (
	grpcPort int
	grpcCmd  = &cobra.Command{
		Use:   "grpc",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Listening GRPC on %d\n", grpcPort)
			listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
			if err != nil {
				panic(err)
			}
			grpcServer := grpc.NewServer()
			pb.RegisterDamageServiceServer(grpcServer, newServer())
			grpcServer.Serve(listen)
		},
	}
)

func init() {
	grpcCmd.Flags().IntVar(&grpcPort, "port", 9090, "GRPC server listen")
	rootCmd.AddCommand(grpcCmd)
}
