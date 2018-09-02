package cmd

import (
	"fmt"

	"github.com/amir20/eztv-watcher/watcher"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(watcher.ConfigFileTemplate)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
