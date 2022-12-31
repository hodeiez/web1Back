package service

import (
	"hodei/web1/db"
	"hodei/web1/drive"
)

// func UploadTrackTest(folder db.DbBase, fileRef string) string {
// 	fileName := drive.Put(folder.String(), fileRef)

// 	trackDb := db.DbTrack{Title: "theTest", Description: "theDescription", FileRef: fileName}
// 	return db.Put(folder, trackDb)

// }
func AddInfoCard(info db.DbInfo) string {
	if info.ImageRef != "" || len(info.ImageAlbum) != 0 || len(info.TracksRef) != 0 || len(info.AlbumsRef) != 0 {
		// info.ImageRef = AddImage(info.ImageRef)
		// info.ImageAlbum = AddImageAlbum(info.ImageAlbum)
		// info.TracksRef = AddTracks(info.TracksRef)
		info.AlbumsRef = AddAlbums(info.AlbumsRef)
	}

	return db.Put(db.InfoCardsBase, info)
}
func AddTrack(trackInfo db.DbTrack) db.DbTrack {
	file := drive.Put(db.TracksBase, trackInfo.FileRef)
	trackInfo.FileRef = file
	trackInfo.Key = db.Put(db.TracksBase, trackInfo)
	return trackInfo
}
func AddTracks(trackInfos []db.DbTrack) []db.DbTrack {
	for i, t := range trackInfos {
		trackInfos[i] = AddTrack(t)
	}
	return trackInfos
}
func AddImage(imageRef string) string {
	return drive.Put(db.ImageBase, imageRef)
}
func AddImageAlbum(imagesRef []string) []string {
	imagesDriveRef := make([]string, len(imagesRef))
	for _, refs := range imagesRef {
		imagesDriveRef = append(imagesDriveRef, drive.Put(db.ImageBase, refs))
	}
	return imagesDriveRef
}
func AddAlbum(albumInfo db.DbAlbum) db.DbAlbum {
	for i, track := range albumInfo.Tracks {
		albumInfo.Tracks[i] = AddTrack(track)
	}
	albumInfo.Key = db.Put(db.AlbumsBase, albumInfo)
	return albumInfo
}

func AddAlbums(albumsInfos []db.DbAlbum) []db.DbAlbum {
	for i, a := range albumsInfos {
		albumsInfos[i] = AddAlbum(a)
	}
	return albumsInfos
}
