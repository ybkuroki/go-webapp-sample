package main

import (
	"github.com/ybkuroki/go-webapp-sample/common"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/router"
)

func main() {
	repository.InitDB()
	db := repository.GetConnection()

	// TODO: switch the following processing by environment
	common.InitMasterData()

	router := router.Init()
	router.Start(":8080")

	defer db.Close()
}
