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
	Nothing
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
func ToInfoType(infoType string) InfoType {
	switch infoType {
	case "Personal":
		return Personal
	case "Professional":
		return Professional
	case "Creative":
		return Creative
	default:
		return Nothing
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
func (album DbAlbum) ToDTO(tracks []DbTrack) DbAlbumDTO {
	return DbAlbumDTO{Key: album.Key, Tracks: tracks, Year: album.Year, Date: album.Year, Title: album.Title, Description: album.Description}

}
func (info DbInfo) ToDTO(tracks []DbTrack, albums []DbAlbumDTO) DbInfoDTO {
	return DbInfoDTO{Key: info.Key, Year: info.Year, Locale: info.Locale, Title: info.Title, Description: info.Description, ImageRef: info.ImageRef, AlbumsRef: albums, TracksRef: tracks, Date: info.Date, ImageAlbum: info.ImageAlbum, InfoType: info.InfoType}
}
