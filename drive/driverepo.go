package drive

import (
	"bufio"
	"fmt"
	"hodei/web1/db"
	envs "hodei/web1/env"
	"os"
	"strings"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/drive"
)

func repo(driveName db.DbBase) *drive.Drive {
	d, _ := deta.New(deta.WithProjectKey(envs.Get("PROJ_KEY")))
	dr, _ := drive.New(d, driveName.String())
	return dr
}
func Put(folderName db.DbBase, fileRef string) string {
	file, err := os.Open(fileRef)
	if err != nil {
		fmt.Println("failed to open file: ", err)
	}
	defer file.Close()

	uploaded, err := repo(folderName).Put(&drive.PutInput{
		Name: getFileNameFromRef(fileRef), Body: bufio.NewReader(file),
	})
	if err != nil {
		fmt.Println("error uploading: ", err)

	}
	return uploaded
}
func getFileNameFromRef(fileRef string) string {
	splitted := strings.Split(fileRef, "/")
	return splitted[len(splitted)-1]
}
