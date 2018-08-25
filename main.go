package main

import (
	"fmt"
	"github.com/amir20/eztv-watcher/eztv"
)

func main() {
	response, _ := eztv.FetchTorrents("6048596")
	fmt.Printf("%+v\n", response)
}
