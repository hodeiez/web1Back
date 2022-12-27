package db

type DbAlbum struct {
	Key         string   `json:"key"`
	Tracks      []string `json:"tracks"`
	Year        string   `json:"year"`
	Date        string   `json:"date"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}
type DbTrack struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	FileRef     string `json:"fileRef"`
}
type DbInfo struct {
	Key         string   `json:"key"`
	Year        string   `json:"year"`
	Locale      string   `json:"locale"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageRef    string   `json:"imageRef"`
	AlbumsRef   []string `json:"albumsRef"`
	TracksRef   []string `json:"tracksRef"`
	Date        string   `json:"date"`
	ImageAlbum  []string `json:"images"`
}
