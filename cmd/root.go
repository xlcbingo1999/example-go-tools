package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tools-demo",
	Short: "run go-tools-demo project",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("fail to start cobra execute: ", err)
		os.Exit(1)
	}
}
