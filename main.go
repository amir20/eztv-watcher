package main

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/amir20/eztv-watcher/eztv"
	"github.com/spf13/viper"
)

func loadDb(path string, object interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		err = saveDb(path, object)

		if err != nil {
			panic(fmt.Errorf("fatal error creating database: %s", err))
		}

		return nil
	}

	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}

	file.Close()

	return err
}

func saveDb(path string, object interface{}) error {
	file, err := os.Create(path)

	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}

	file.Close()

	return err
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/eztv")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %s", err))
	}

	database := NewDatabase(viper.GetString("database.path"))
	err = database.CreateIfMissing()
	if err != nil {
		panic(fmt.Errorf("fatal error creating database: %s", err))
	}

	err = database.Load()
	if err != nil {
		panic(fmt.Errorf("fatal error loading database: %s", err))
	}

	for _, id := range viper.GetStringSlice("ids") {
		response, _ := eztv.FetchTorrents(id)
		fmt.Printf("%+v\n", response.Torrents[0].Title)
		if lastID, ok := database.GetValue(id); ok {
			println(lastID)
		}

		database.UpdateValue(id, response.Torrents[0].ID)
	}

	err = database.Save()
	if err != nil {
		panic(fmt.Errorf("fatal error writing database: %s", err))
	}

}
