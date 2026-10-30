package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
