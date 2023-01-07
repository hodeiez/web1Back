package main

import (
	// "hodei/web1/db"

	"hodei/web1/controller"
	gui "hodei/web1/mygui"

	// "hodei/web1/service"
	"os"
)

func main() {
	//test db
	if len(os.Args) > 1 && os.Args[1] == "local" {
		println("running local setup")
		gui.Init()
	} else {
		controller.InitApi()
		println("buuuu")
	}
	// albumKey := "ngshzau1u2nk"
	// infoCardKey := "bkf482hx7vmx"
	// trackKey := "23lob0l15zk8"
	// jason, _ := json.Marshal(service.GetAlbumDTOByKey(albumKey))
	// fmt.Println(string(jason))
	// infojson, _ := json.MarshalIndent(service.GetInfoCardDTOByKey(infoCardKey), "", "   ")
	// fmt.Println(string(infojson))
	// os.WriteFile("file.mp3", service.GetAudioFileByKey(trackKey), os.ModeDevice.Perm())
	// infoCardByType, _ := json.MarshalIndent(db.GetInfoCardsByType(db.Creative), "", "   ")
	// fmt.Println(string(infoCardByType))
}
