package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/amir20/eztv-watcher/eztv"
	"github.com/amir20/eztv-watcher/watcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(syncCmd)
}

// initCmd represents the init command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs bittorrents by fetching each show and downloading new shows.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("config.yml not found.\n%s", err)
		}

		database := watcher.NewDatabase(viper.GetString("database.path"))
		err := database.CreateIfMissing()
		if err != nil {
			log.Fatalf("Could not write database file. Error: %s", err)
		}

		err = database.Load()
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}

		for _, id := range viper.GetStringSlice("ids") {
			log.Printf("Checking [%s].\n", id)
			response, err := eztv.FetchTorrents(id)

			if err != nil {
				log.Fatalf("Error fetching torrents from EZTV: %s", err)
			}

			var lastUpdated time.Time
			if v, ok := database.GetValue(id); ok {
				lastUpdated = time.Unix(int64(v), 0)
			} else {
				lastUpdated = time.Now().AddDate(0, 0, -7)
			}

			for _, torrent := range response.Torrents {
				released := time.Unix(int64(torrent.DateReleasedUnix), 0)
				if released.After(lastUpdated) && isNotBlacklisted(torrent) && isWhitelisted(torrent) {
					log.Printf("Found a new torrent [%s].\n", torrent.Title)
					f := filepath.Join(viper.GetString("torrent_watch_dir"), torrent.Filename+".torrent")
					if err := watcher.DownloadFile(f, torrent.TorrentURL); err != nil {
						log.Fatalf("Could not write torrent file [%s]. Error %s", f, err)
					}
					fmt.Printf("Downloading a new episode [%s].\n", torrent.Title)
				}
			}

			if len(response.Torrents) > 0 {
				database.UpdateValue(id, response.Torrents[0].DateReleasedUnix)
			}
		}

		if err := database.Save(); err != nil {
			log.Fatalf("Could not save database file. Error %s", err)
		}
	},
}

func isNotBlacklisted(torrent eztv.Torrent) bool {
	for _, match := range viper.GetStringSlice("matches.blacklist") {
		if strings.Contains(torrent.Title, match) {
			return false
		}
	}
	return true
}

func isWhitelisted(torrent eztv.Torrent) bool {
	for _, match := range viper.GetStringSlice("matches.whitelist") {
		if strings.Contains(torrent.Title, match) {
			return true
		}
	}

	return false
}
