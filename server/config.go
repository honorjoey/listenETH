package server

import (
	"github.com/joeylaker/loadConfig"
	"log"
	"strconv"
)

func initConfig() {
	config := loadConfig.NewConfig()
	config.LoadConfig("./config.ini")
	Net = config.GetConfig("ETHNode")
	start := config.GetConfig("startBlockNumber")
	number, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		log.Println("config blockNumber is not a number")
	}
	startBlockNumber = number
	addressListFilePath = config.GetConfig("addressListFile")
	tokenListFilePath = config.GetConfig("tokenListFile")
	//Email
	EmailUser = config.GetConfig("EmailUser")
	EmailPassword = config.GetConfig("EmailPassword")
	EmailHost = config.GetConfig("EmailHost")
	EmailTo = config.GetConfig("EmailTo")
}
