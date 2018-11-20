package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/amir20/eztv-watcher/watcher"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes config file for the first time",
	Run: func(cmd *cobra.Command, args []string) {
		path := os.ExpandEnv("/home/$USER/.config/eztv/config.yml")

		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			log.Fatalf("Cannot create directory: %s\n%s", filepath.Dir(path), err)
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			file, err := os.Create(path)
			if err == nil {
				defer file.Close()
				if _, err := file.Write(watcher.ConfigFileTemplate); err != nil {
					log.Fatalf("Could not write %s\n%s", path, err)
				}
				log.Printf("Successfully created %s", path)
			} else {
				log.Fatalf("Could not create %s\n%s", path, err)
			}
		} else {
			log.Fatalf("%s already exists. Aborting.", path)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
