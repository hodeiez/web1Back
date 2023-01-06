package service

import (
	"hodei/web1/db"
	"hodei/web1/drive"
)

//*************ADD**********************//
func AddInfoCard(info db.DbInfoDTO) string {

	info.ImageRef = AddImage(info.ImageRef)
	info.ImageAlbum = AddImageAlbum(info.ImageAlbum)
	info.TracksRef = AddTracks(info.TracksRef)
	info.AlbumsRef = AddAlbums(info.AlbumsRef)

	return db.Put(db.InfoCardsBase, DbInfoConverter(info))
}
func AddTrack(trackInfo db.DbTrack) db.DbTrack {
	file := drive.Put(db.TracksBase, trackInfo.FileRef)
	trackInfo.FileRef = file
	trackInfo.Key = db.Put(db.TracksBase, trackInfo)
	return trackInfo
}
func AddTracks(trackInfos []db.DbTrack) []db.DbTrack {
	if len(trackInfos) != 0 {
		for i, t := range trackInfos {
			trackInfos[i] = AddTrack(t)
		}
	}
	return trackInfos
}
func AddImage(imageRef string) string {
	if imageRef != "" {
		return drive.Put(db.ImageBase, imageRef)
	}
	return imageRef
}
func AddImageAlbum(imagesRef []string) []string {
	if len(imagesRef) != 0 {
		for i, refs := range imagesRef {
			imagesRef[i] = AddImage(refs)
		}
	}
	return imagesRef

}
func AddAlbum(albumInfo db.DbAlbumDTO) db.DbAlbumDTO {

	for i, track := range albumInfo.Tracks {
		albumInfo.Tracks[i] = AddTrack(track)
	}
	converted := DbAlbumConverter(albumInfo)
	albumInfo.Key = db.Put(db.AlbumsBase, converted)

	return albumInfo
}

func AddAlbums(albumsInfos []db.DbAlbumDTO) []db.DbAlbumDTO {
	if len(albumsInfos) != 0 {
		for i, a := range albumsInfos {
			albumsInfos[i] = AddAlbum(a)
		}
	}
	return albumsInfos
}

//**************************************UPDATES****************//
func UpdateTrack(trackInfo db.DbTrack) {
	db.Update(db.TracksBase, trackInfo.Key, trackInfo)
}

//*******************************GET***************************//
func GetAlbumDTOByKey(key string) db.DbAlbumDTO {
	album := db.GetAlbumByKey(key)
	tracks := make([]db.DbTrack, len(album.Tracks))
	for i, t := range album.Tracks {
		tracks[i] = db.GetTrackByKey(t)
	}
	return album.ToDTO(tracks)
}
func GetAlbumsDTOByKey(keys []string) []db.DbAlbumDTO {
	albums := make([]db.DbAlbumDTO, len(keys))
	for i, k := range keys {
		albums[i] = GetAlbumDTOByKey(k)
	}
	return albums
}
func GetTracksByKey(keys []string) []db.DbTrack {
	tracks := make([]db.DbTrack, len(keys))
	for i, k := range keys {
		tracks[i] = db.GetTrackByKey(k)
	}
	return tracks
}
func GetInfoCardDTOByKey(key string) db.DbInfoDTO {
	info := db.GetInfoCardByKey(key)
	return info.ToDTO(GetTracksByKey(info.TracksRef), GetAlbumsDTOByKey(info.AlbumsRef))
}
func GetInfoCardDTOByType(info db.InfoType) []db.DbInfoDTO {
	infos := db.GetInfoCardsByType(info)
	return infosToDTO(infos)
}
func GetInfoCardDTOByRange(from string, to string) []db.DbInfoDTO {
	infos := db.GetInfoCardsByRange(from, to)
	return infosToDTO(infos)
}
func GetAudioFileByRef(ref string) []byte {
	return drive.GetAudioFile(ref)
}
func GetAudioFileByKey(key string) []byte {
	return drive.GetAudioFile(db.GetTrackByKey(key).FileRef)
}
