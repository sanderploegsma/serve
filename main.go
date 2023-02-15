package main

import (
	"fmt"

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
		fmt.Printf("Serving %s on port %d", args[0], port)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 0, "Port to serve on. A random free port will be used automatically if not provided")
}

func main() {
	rootCmd.Execute()
}
