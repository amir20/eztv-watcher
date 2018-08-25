package main

import (
	"encoding/gob"
	"os"
)

// Database represent a map on disk
type Database interface {
	CreateIfMissing() error
	Load() error
	Save() error
	UpdateValue(string, int)
	GetValue(string) (int, bool)
}

type database struct {
	path string
	data map[string]int
}

// NewDatabase creates a new instance of in memory database
func NewDatabase(path string) Database {
	return &database{path, make(map[string]int)}
}

func (db *database) CreateIfMissing() error {
	if _, err := os.Stat(db.path); os.IsNotExist(err) {
		return db.Save()
	}

	return nil
}

func (db *database) Load() error {
	file, err := os.Open(db.path)

	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&db.data)
	}

	file.Close()

	return err
}

func (db *database) Save() error {
	file, err := os.Create(db.path)

	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(&db.data)
	}

	file.Close()

	return err
}

func (db *database) UpdateValue(k string, v int) {
	db.data[k] = v
}

func (db *database) GetValue(k string) (int, bool) {
	v, ok := db.data[k]
	return v, ok
}
