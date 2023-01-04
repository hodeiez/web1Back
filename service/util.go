package service

import "hodei/web1/db"

func DbInfoConverter(info db.DbInfoDTO) db.DbInfo {
	return db.DbInfo{Key: info.Key,
		Year:        info.Year,
		Locale:      info.Locale,
		Title:       info.Title,
		Description: info.Description,
		ImageRef:    info.ImageRef,
		AlbumsRef:   AlbumsToRef(info.AlbumsRef),
		TracksRef:   TracksToRef(info.TracksRef),
		Date:        info.Date,
		ImageAlbum:  info.ImageAlbum,
		InfoType:    info.InfoType}
}
func DbAlbumConverter(album db.DbAlbumDTO) db.DbAlbum {
	return db.DbAlbum{Key: album.Key, Tracks: TracksToRef(album.Tracks), Year: album.Year, Date: album.Date, Title: album.Title, Description: album.Description}
}
func TracksToRef(tracks []db.DbTrack) []string {
	trackRefs := make([]string, len(tracks))
	if len(tracks) != 0 {
		for i, t := range tracks {
			trackRefs[i] = t.Key
		}
	}
	return trackRefs
}
func AlbumsToRef(albums []db.DbAlbumDTO) []string {
	albumsRefs := make([]string, len(albums))
	if len(albums) != 0 {
		for i, a := range albums {
			albumsRefs[i] = a.Key
		}
	}
	return albumsRefs
}
func infosToDTO(infos []db.DbInfo) []db.DbInfoDTO {
	infosDTO := make([]db.DbInfoDTO, len(infos))
	for i, inf := range infos {
		infosDTO[i] = inf.ToDTO(GetTracksByKey(inf.TracksRef), GetAlbumsDTOByKey(inf.AlbumsRef))
	}
	return infosDTO
}
