package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetch(url string) (*EztvResponse, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if r.StatusCode != 200 {
		return nil, fmt.Errorf("none 200 response from [%s]", url)
	}

	response := &EztvResponse{}
	if err = json.NewDecoder(r.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	response, _ := fetch("https://eztv.ag/api/get-torrents?imdb_id=6048596")
	fmt.Printf("%+v\n", response)
}
