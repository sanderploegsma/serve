package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	version = "dev"

	port int
)

var rootCmd = &cobra.Command{
	Use:     "serve path/to/serve",
	Short:   "Serve helps you serve static content",
	Long:    "A static content server for your command line, built with Go.",
	Version: version,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		options := ServeOptions{Port: port}

		if err := serve(path, options); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 0, "Port to serve on. A random free port will be used automatically if not provided")
}

func main() {
	rootCmd.Execute()
}

type ServeOptions struct {
	Port int
}

func serve(path string, options ServeOptions) error {
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", options.Port))
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	log.Printf("Serving %s on http://localhost:%d\n", path, listener.Addr().(*net.TCPAddr).Port)
	return http.Serve(listener, nil)
}
