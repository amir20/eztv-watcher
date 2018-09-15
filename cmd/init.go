package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
