package main

import (
	kernel2 "go-api/app/commands/kernel"
	"os"
)

func main()  {
	if err := kernel2.MainCmd.Execute(); err != nil {
		os.Exit(1)
	}
}