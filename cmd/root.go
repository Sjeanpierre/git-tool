package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GitTool",
	Short: "Manage PR interactions from the commandline",
	Long: `Simple management utility for Github PRs`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}