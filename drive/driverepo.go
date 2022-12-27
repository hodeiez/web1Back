package drive

import (
	"bufio"
	"fmt"
	envs "hodei/web1/env"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/drive"
)

func repo() *drive.Drive {
	d, _ := deta.New(deta.WithProjectKey(envs.Get("PROJ_KEY")))
	dr, _ := drive.New(d, "web1")
	return dr
}
func Put(fileName string, fileRef string) string {
	file, err := os.Open(fileRef)
	if err != nil {
		fmt.Println("failed to open file: ", err)
	}
	defer file.Close()

	uploaded, err := repo().Put(&drive.PutInput{
		Name: fileName, Body: bufio.NewReader(file),
	})
	if err != nil {
		fmt.Println("error uploading: ", err)

	}
	return uploaded
}
