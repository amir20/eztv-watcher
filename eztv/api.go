package eztv

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func FetchTorrents(id string) (*EztvResponse, error) {
	r, err := http.Get(fmt.Sprintf("https://eztv.ag/api/get-torrents?imdb_id=%s", id))
	if err != nil {
		return nil, err
	} else if r.StatusCode != 200 {
		return nil, fmt.Errorf("none 200 response from [%s]", id)
	}

	response := &EztvResponse{}
	if err = json.NewDecoder(r.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
