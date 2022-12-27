package db

import (
	"fmt"
	envs "hodei/web1/env"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

func repo() *base.Base {
	d, _ := deta.New(deta.WithProjectKey(envs.Get("PROJ_KEY")))
	db, _ := base.New(d, "web1")
	return db

}

func Put(info interface{}) string {
	inserted, err := repo().Insert(info)
	if err != nil {
		fmt.Println("failed to insert: ", err)

	}
	return inserted
}
