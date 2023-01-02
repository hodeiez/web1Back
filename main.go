package main

import (
	// "hodei/web1/db"
	"fmt"
	"hodei/web1/service"
	"os"
)

func main() {
	//test db
	if len(os.Args) > 1 && os.Args[1] == "local" {
		println("running local setup")
	} else {
		println("buuuu")
	}
	fmt.Println(service.GetAlbumDTOByKey("ngshzau1u2nk"))
	// track1 := db.DbTrack{Title: "one", Description: "desc", Date: "2022-12-32", FileRef: "./test/audio.mp3"}
	// track2 := db.DbTrack{Title: "two", Description: "desc", Date: "2022-12-32", FileRef: "./test/audio2.mp3"}
	// album := db.DbAlbumDTO{Tracks: []db.DbTrack{track1, track2}, Year: "2022", Date: "2022-21-21", Title: "title", Description: "desc"}
	// info := db.DbInfoDTO{Year: "2022", Locale: "eng", Title: "a", Description: "desc", AlbumsRef: []db.DbAlbumDTO{album}, Date: "2022", InfoType: db.Creative.String()}

	// println(service.AddInfoCard(info))
}
