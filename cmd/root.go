package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gosls",
	Short: "A CLI to create and manage Serverless applications for AWS Lambda",
	Long: `Go Serverless is a CLI to create and manage applications with the Serverless Framework.
It is a wrapper around the Serverless Framework CLI and provides a simple way to
create and manage Serverless applications using the Go programming language.`,
	Version: "0.0.1",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
