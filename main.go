package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	versionOption = flag.Bool("v", false, "Prints the version")
)

func main() {
	flag.Parse()

	if *versionOption {
		fmt.Printf("Serve %s\nCommit: %s\nDate: %s", version, commit, date)
		os.Exit(0)
	}
}
