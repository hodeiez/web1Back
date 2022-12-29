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
	if info.ImageRef != "" {

	} else if len(info.ImageAlbum) != 0 {

	} else if len(info.TracksRef) != 0 {

	} else if len(info.TracksRef) != 0 {

	}

	return db.Put(db.InfoCardsBase, info)
}
func AddTrack(trackInfo db.DbTrack) string {
	file := drive.Put(db.TracksBase, trackInfo.FileRef)
	trackInfo.FileRef = file
	return db.Put(db.TracksBase, trackInfo)
}
func AddImage(imageRef string) string {
	return drive.Put(db.ImageBase, imageRef)
}
func AddAlbum(albumInfo db.DbAlbum) string {
	for _, track := range albumInfo.Tracks {
		track.FileRef = AddTrack(track)
	}
	return db.Put(db.AlbumsBase, albumInfo)
}
