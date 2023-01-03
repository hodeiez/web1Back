package main

import (
	// "hodei/web1/db"
	"encoding/json"
	"fmt"
	"hodei/web1/db"
	// "hodei/web1/service"
	"os"
)

func main() {
	//test db
	if len(os.Args) > 1 && os.Args[1] == "local" {
		println("running local setup")
	} else {
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
	infoCardByType, _ := json.MarshalIndent(db.GetInfoCardByType(db.Creative), "", "   ")
	fmt.Println(string(infoCardByType))

}
