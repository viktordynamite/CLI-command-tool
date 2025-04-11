// Linker file

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/viktordynamite/CLI-command-tool/internal/storage"
)

func init() {
	var execCmd = &cobra.Command{
		Use:   "exec [index]",
		Short: "Execute a command",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := 0
			fmt.Sscanf(args[0], "%d", &index)

			if index < 1 || index > len(storage.Commands) {
				return fmt.Errorf("invalid command index: %d (valid range: 1-%d)", index, len(storage.Commands))
			}

			command := storage.Commands[index-1].Command
			fmt.Printf("Executing: %s\n", command)

			// On Windows, use cmd.exe instead of sh
			execCmd := exec.Command("cmd", "/C", command)
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr
			execCmd.Stdin = os.Stdin

			return execCmd.Run()
		},
	}

	rootCmd.AddCommand(execCmd)
}
