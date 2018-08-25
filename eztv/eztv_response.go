package eztv

// Response represent a response from EZTV
type Response struct {
	ImdbID        string    `json:"imdb_id"`
	TorrentsCount int       `json:"torrents_count"`
	Limit         int       `json:"limit"`
	Page          int       `json:"page"`
	Torrents      []Torrent `json:"torrents"`
}

// Torrent is one torrent in a response
type Torrent struct {
	ID               int    `json:"id"`
	Hash             string `json:"hash"`
	Filename         string `json:"filename"`
	EpisodeURL       string `json:"episode_url"`
	TorrentURL       string `json:"torrent_url"`
	MagnetURL        string `json:"magnet_url"`
	Title            string `json:"title"`
	ImdbID           string `json:"imdb_id"`
	Season           string `json:"season"`
	Episode          string `json:"episode"`
	SmallScreenshot  string `json:"small_screenshot"`
	LargeScreenshot  string `json:"large_screenshot"`
	Seeds            int    `json:"seeds"`
	Peers            int    `json:"peers"`
	DateReleasedUnix int    `json:"date_released_unix"`
	SizeBytes        string `json:"size_bytes"`
}
