package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xlcbingo1999/example-go-tools-xlc/ginDemo"
)

var ginDemoCmd = &cobra.Command{
	Use:   "gin_demo",
	Short: "Run gin_demo",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln("Recover err", err)
			}
		}()

		ginDemo.RunGinDemo()
	},
}

func init() {
	ginDemoCmd.Flags().StringVarP(&ginDemo.ServePort, "gin-serve-port", "", "45677", "mutex demo serve port")
	rootCmd.AddCommand(ginDemoCmd)
}
