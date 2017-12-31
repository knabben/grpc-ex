package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/knabben/grpc-ex/damage"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	remoteGrpc string
	port       int

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
)

func init() {
	serveCmd.Flags().StringVar(&remoteGrpc, "grpc", "localhost:9090", "Remote GRPC address")
	serveCmd.Flags().IntVar(&port, "port", 8080, "Local port")
	rootCmd.AddCommand(serveCmd)
}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	fmt.Println("preflight request for %s", r.URL.Path)
	return
}

func serve() error {
	fmt.Printf("Access remote GRPC on %s\n", remoteGrpc)
	fmt.Printf("Listening on port %d\n", port)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterDamageServiceHandlerFromEndpoint(
		ctx, gwmux, remoteGrpc, dialOpts)
	if err != nil {
		panic(err)
	}

	mux.Handle("/", gwmux)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), allowCORS(mux))
}
