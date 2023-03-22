package db

import (
	"fmt"
	envs "hodei/web1/env"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

func repo(baseName DbBase) *base.Base {
	d, _ := deta.New(deta.WithProjectKey(envs.Get("PROJ_KEY")))
	db, _ := base.New(d, baseName.String())
	return db

}

func Put(baseName DbBase, info interface{}) string {
	inserted, err := repo(baseName).Insert(info)
	if err != nil {
		fmt.Println("failed to insert: ", err)

	}
	return inserted
}
func Update(baseName DbBase, key string, info interface{}) {
	err := repo(baseName).Update(key, info.(base.Updates))
	if err != nil {
		fmt.Println("failed to update: ", err)
	}
}
func GetByKey(baseName DbBase, key string, dest interface{}) error {
	err := repo(baseName).Get(key, dest)
	if err != nil {
		fmt.Println("failed to insert: ", err)

	}
	return err
}
func GetTrackByKey(key string) (DbTrack, error) {
	dest := DbTrack{}
	err := repo(TracksBase).Get(key, &dest)
	if err != nil {
		fmt.Println("failed to get: ", err)

	}
	return dest, err
}
func GetAlbumByKey(key string) DbAlbum {
	dest := DbAlbum{}
	err := repo(AlbumsBase).Get(key, &dest)
	if err != nil {
		fmt.Println("failed to insert: ", err)

	}
	return dest
}
func GetInfoCardByKey(key string) DbInfo {
	dest := DbInfo{}
	err := repo(InfoCardsBase).Get(key, &dest)
	if err != nil {
		fmt.Println("failed to insert: ", err)

	}
	return dest
}
func GetInfoCardsByType(info InfoType) []DbInfo {
	dest := []DbInfo{}
	_, err := repo(InfoCardsBase).Fetch(&base.FetchInput{Q: base.Query{{"type": info.String()}}, Dest: &dest})
	if err != nil {
		fmt.Println("couldn't fetch infocards by type: ", err)
	}
	return dest
}
func GetInfoCardsByRange(from string, to string) []DbInfo {
	dest := []DbInfo{}
	_, err := repo(InfoCardsBase).Fetch(&base.FetchInput{Q: base.Query{{"date?r": []string{from, to}}}, Dest: &dest})
	if err != nil {
		fmt.Println("couldn't fetch infocards by range:", err)
	}
	return dest

}
func GetAllTracks() []DbTrack {
	dest := []DbTrack{}
	_, err := repo(TracksBase).Fetch(&base.FetchInput{Dest: &dest})
	if err != nil {
		fmt.Println("couldn't fetch all tracks:", err)
	}
	return dest

}
func GetAllAlbums() []DbAlbum {
	dest := []DbAlbum{}
	_, err := repo(AlbumsBase).Fetch(&base.FetchInput{Dest: &dest})
	if err != nil {
		fmt.Println("couldn't fetch all tracks:", err)
	}
	return dest

}
func GetInfoCardsByLocale(locale string) []DbInfo {
	dest := []DbInfo{}
	_, err := repo(InfoCardsBase).Fetch(&base.FetchInput{Q: base.Query{{"locale": locale}}, Dest: &dest})
	if err != nil {
		fmt.Println("couldn't fetch infocards by locale: ", err)
	}
	return dest
}
