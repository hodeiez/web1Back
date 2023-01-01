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
