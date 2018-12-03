package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "eztv-watcher",
	Short: "Small CLI for fetching and syncing TV show bit torrents from EZTV.",
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
	if path, ok := os.LookupEnv("SNAP_USER_COMMON"); ok {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath("$HOME/.config/eztv")
	viper.AddConfigPath(".")
	if _, ok := os.LookupEnv("SNAP_USER_COMMON"); ok {
		viper.SetDefault("database.path", os.ExpandEnv("$SNAP_USER_COMMON/db.bin"))
	} else {
		viper.SetDefault("database.path", os.ExpandEnv("$HOME/.config/eztv/db.bin"))
	}

	viper.SetDefault("matches.whitelist", []string{})
	viper.SetDefault("matches.blacklist", []string{})
}
