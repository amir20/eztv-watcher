package main

import (
	"fmt"
	"time"

	"github.com/amir20/eztv-watcher/eztv"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/eztv")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
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
		response, _ := eztv.FetchTorrents(id)

		var lastUpdated time.Time
		if v, ok := database.GetValue(id); ok {
			lastUpdated = time.Unix(int64(v), 0)
		} else {
			lastUpdated = time.Now().AddDate(0, 0, -7)
		}
		println(lastUpdated.String())

		database.UpdateValue(id, response.Torrents[0].DateReleasedUnix)
	}

	err = database.Save()
	if err != nil {
		panic(fmt.Errorf("fatal error writing database: %s", err))
	}

}
