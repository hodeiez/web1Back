package main

import (
	"hodei/web1/db"
	"hodei/web1/service"
)

func main() {
	//test db

	// info := db.DbInfo{Year: "2022",
	// 	Locale:      "eus",
	// 	Title:       "Title",
	// 	Description: "deskribapena",
	// 	Date:        "2022-20-11",
	// 	AlbumsRef:   []string{"fdsa", "fdsa", "gdfsgfds"},
	// }
	// fmt.Println(info.ImageRef == "")
	// fmt.Println(len(info.ImageAlbum) == 0)
	// println(db.Put(db.InfoCardsBase, info))
	// //test drive
	// println(drive.Put("atest", "./test/test.jpg"))

	// println(service.UploadTrackTest(db.TracksBase, "./test/test.jpg"))
	track1 := db.DbTrack{Title: "one", Description: "desc", Date: "2022-12-32", FileRef: "./test/audio.mp3"}
	track2 := db.DbTrack{Title: "two", Description: "desc", Date: "2022-12-32", FileRef: "./test/audio2.mp3"}
	album := db.DbAlbum{Tracks: []db.DbTrack{track1, track2}, Year: "2022", Date: "2022-21-21", Title: "title", Description: "desc"}
	// println(service.AddAlbum(album).Title)
	info := db.DbInfo{Year: "2022", Locale: "eng", Title: "a", Description: "desc", AlbumsRef: []db.DbAlbum{album}, Date: "2022", InfoType: db.Creative.String()}

	println(service.AddInfoCard(info))
}
