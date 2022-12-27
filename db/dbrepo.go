package db

import (
	envs "hodei/web1/env"
)

func GetThis() string {

	return envs.Get("BAT")
}
