package cmd

import (
	pb "github.com/knabben/grpc/damage"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

// beerCmd represents the beer command
var beerCmd = &cobra.Command{
	Use:   "beer",
	Short: "smoke",
	Run: func(cmd *cobra.Command, args []string) {
		BeerServe()
	},
}

type server struct{}

func newServer() pb.DamageServiceServer {
	return new(server)
}

func (s *server) Damage(ctx context.Context, msg *pb.DamageMessage) (*pb.DamageMessage, error) {
	return &pb.DamageMessage{Value: msg.Value}, nil
}

func BeerServe() error {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDamageServiceServer(grpcServer, newServer())
	return grpcServer.Serve(listen)
}

func init() {
	rootCmd.AddCommand(beerCmd)
}
