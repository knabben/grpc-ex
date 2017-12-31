package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/knabben/grpc-ex/damage"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

type server struct{}

func newServer() pb.DamageServiceServer {
	return new(server)
}

func (s *server) Damage(ctx context.Context, msg *pb.DamageMessage) (*pb.DamageMessage, error) {
	return &pb.DamageMessage{Value: msg.Value}, nil
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			fmt.Println("grpc")
			grpcServer.ServeHTTP(w, r)
		} else {
			fmt.Println("http")
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func serve() error {
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000"))}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDamageServiceServer(grpcServer, newServer())
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: demoAddr,
		RootCAs:    demoCertPool,
	})
	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()

	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	err := pb.RegisterDamageServiceHandlerFromEndpoint(
		ctx, gwmux, demoAddr, dopts)
	if err != nil {
		fmt.Printf("serve %v\n", err)
		return err
	}

	mux.Handle("/", gwmux)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK\n")
	})

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    demoAddr,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*demoKeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	fmt.Printf("grpc on port: %d\n", port)
	if err = srv.Serve(tls.NewListener(conn, srv.TLSConfig)); err != nil {
		fmt.Println("ListenAndServe: ", err)
		return nil
	}
	return nil
}
