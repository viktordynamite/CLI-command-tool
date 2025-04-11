// Root and description file

package main

import (
	"github.com/spf13/cobra"
	"github.com/viktordynamite/CLI-command-tool/internal/storage"
)

var rootCmd = &cobra.Command{
	Use:   "cli-tool",
	Short: "A tool to manage your command line commands",
	Long:  "A tool that helps you store, search, and retrieve frequently used command line commands so you don't have to remember them all.",
}

// Execute executes the root command
func Execute() error {
	if err := storage.LoadCommands(); err != nil {
		return err
	}

	return rootCmd.Execute()
}

func init() {
	// Define flags and configuration settings here
}
