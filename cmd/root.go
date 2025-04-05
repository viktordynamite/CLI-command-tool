// Defines root command and setsup Cobra

package cmd

import (
	// "fmt"
	// "os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "TCM-tool",
	Short: "A tool to manage your command line commands",
	Long:  "A tool that helps you store, search, and retrieve frequently used command line commands so you don't have to remember them all.",
}

// executes all child commands to root command and sets flags
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// define flags and configuration settings
	// called by main.main()
}
