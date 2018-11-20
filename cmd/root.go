package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "eztv-watcher",
	Short: "Small CLI for fetching and synching TV show bittorrents from EZTV.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/eztv")
	viper.AddConfigPath(os.ExpandEnv("/home/$USER/.config/eztv"))
	viper.AddConfigPath("$HOME/.config/eztv")
	viper.AddConfigPath(".")
	viper.SetDefault("database.path", os.ExpandEnv("/home/$USER/.config/eztv/db.bin"))
	viper.SetDefault("matches.whitelist", []string{})
	viper.SetDefault("matches.blacklist", []string{})
}
