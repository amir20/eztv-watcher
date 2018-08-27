package eztv

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchTorrents returns a list of torrent for a given IMDB 
func FetchTorrents(id string) (*Response, error) {
	r, err := http.Get(fmt.Sprintf("https://eztv.ag/api/get-torrents?imdb_id=%s", id))
	if err != nil {
		return nil, err
	} else if r.StatusCode != 200 {
		return nil, fmt.Errorf("none 200 response from [%s]", id)
	}

	response := &Response{}
	if err = json.NewDecoder(r.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
