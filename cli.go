package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArgs() *string {
	var domain = flag.String("domain", "", "The domain you want to modify.")
	flag.Parse()
	if len(*domain) == 0 {
		fmt.Println("invalid domain")
		os.Exit(1)
	}
	return domain
}
