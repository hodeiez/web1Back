package service

import (
	"hodei/web1/db"
	"hodei/web1/drive"
)

func UploadTrack(name string, fileRef string) string {
	fileName := drive.Put(name, fileRef)

	trackDb := db.DbTrack{Title: "theTest", Description: "theDescription", FileRef: fileName}
	return db.Put(trackDb)

}
