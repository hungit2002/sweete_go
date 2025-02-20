package main

import (
	"flag"
	"sweete/user_service/cmd"
	"sweete/user_service/config"
)

func main() {
	flag.Parse()
	_ = config.GetInstance()
	rootCmd := cmd.Root()
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
