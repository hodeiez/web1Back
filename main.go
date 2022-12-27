package main

import (
	db "hodei/web1/db"
	"hodei/web1/drive"
	"hodei/web1/service"
)

func main() {
	//test db

	info := db.DbInfo{Year: "2022",
		Locale:      "eus",
		Title:       "Title",
		Description: "deskribapena",
		Date:        "2022-20-11",
		AlbumsRef:   []string{"fdsa", "fdsa", "gdfsgfds"},
	}

	println(db.Put(info))
	//test drive
	println(drive.Put("atest", "./test/test.jpg"))

	println(service.UploadTrack("atest", "./test/test.jpg"))
}
