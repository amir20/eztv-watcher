package main

import "github.com/amir20/eztv-watcher/cmd"

//go:generate go run gen.go

func main() {
	cmd.Execute()
}
