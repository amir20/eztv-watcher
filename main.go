package main

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/amir20/eztv-watcher/eztv"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/eztv")
	viper.AddConfigPath(".")

	usr, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("cannot find user's home directory: %s", err))
	}

	viper.SetDefault("database.path", filepath.Join(usr.HomeDir, ".config/eztv/db.bin"))
	viper.SetDefault("matches.whitelist", []string{})
	viper.SetDefault("matches.blacklist", []string{})

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}
}

func main() {
	database := NewDatabase(viper.GetString("database.path"))
	err := database.CreateIfMissing()
	if err != nil {
		panic(fmt.Errorf("fatal error creating database: %s", err))
	}

	err = database.Load()
	if err != nil {
		panic(fmt.Errorf("fatal error loading database: %s", err))
	}

	for _, id := range viper.GetStringSlice("ids") {
		fmt.Printf("Checking [%s].\n", id)
		response, _ := eztv.FetchTorrents(id)

		var lastUpdated time.Time
		if v, ok := database.GetValue(id); ok {
			lastUpdated = time.Unix(int64(v), 0)
		} else {
			lastUpdated = time.Now().AddDate(0, 0, -7)
		}

		for _, torrent := range response.Torrents {
			released := time.Unix(int64(torrent.DateReleasedUnix), 0)
			if released.After(lastUpdated) && isNotBlacklisted(torrent) && isWhitelisted(torrent) {
				fmt.Printf("Found a new torrent [%s].\n", torrent.Title)
				f := filepath.Join(viper.GetString("torrent_watch_dir"), torrent.Filename+".torrent")
				if err := DownloadFile(f, torrent.TorrentURL); err != nil {
					panic(fmt.Errorf("fatal error when writing torrent file [%s] with error %s", f, err))
				}
			}
		}

		if len(response.Torrents) > 0 {
			database.UpdateValue(id, response.Torrents[0].DateReleasedUnix)
		}
	}

	err = database.Save()
	if err != nil {
		panic(fmt.Errorf("fatal error writing database: %s", err))
	}
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
