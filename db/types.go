package db

type DbAlbumDTO struct {
	Key         string    `json:"key"`
	Tracks      []DbTrack `json:"tracks"`
	Year        string    `json:"year"`
	Date        string    `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
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
type DbInfoDTO struct {
	Key         string       `json:"key"`
	Year        string       `json:"year"`
	Locale      string       `json:"locale"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ImageRef    string       `json:"imageRef"`
	AlbumsRef   []DbAlbumDTO `json:"albumsRef"`
	TracksRef   []DbTrack    `json:"tracksRef"`
	Date        string       `json:"date"`
	ImageAlbum  []string     `json:"images"`
	InfoType    string       `json:"type"`
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
	InfoType    string   `json:"type"`
}
type InfoType int

const (
	Personal InfoType = iota
	Professional
	Creative
)

func (infoType InfoType) String() string {
	switch infoType {
	case Personal:
		return "Personal"
	case Professional:
		return "Professional"
	case Creative:
		return "Creative"
	default:
		return "-"
	}
}

type DbBase int

const (
	TracksBase DbBase = iota
	AlbumsBase
	InfoCardsBase
	ImageBase
)

func (dbBase DbBase) String() string {
	switch dbBase {
	case TracksBase:
		return "Tracks"
	case AlbumsBase:
		return "Albums"
	case InfoCardsBase:
		return "InfoCards"
	case ImageBase:
		return "Images"
	default:
		return "-"
	}
}
