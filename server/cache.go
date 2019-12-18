package server

import (
	"log"
	"path/filepath"
)

func InitCache() {
	//读取配置信息
	initConfig()
	//读取要监听的地址列表
	ReadAddressList()
	//读取要监听的代币地址列表
	ReadTokenList()
	//获取最新区块高度
	go getLatestBlockNumber()
	//同步区块
	go getTxs()
	//定时发送邮件
	go sendEThMessage()
}

func ReadAddressList() {
	filePath, _ := filepath.Abs(addressListFilePath)
	addresses, err := ReadFile(filePath)
	if err != nil {
		log.Println("read the address list failed!")
	}
	addressListLock.Lock()
	addressList = *addresses
	addressListLock.Unlock()
}

func ReadTokenList() {
	filePath, _ := filepath.Abs(tokenListFilePath)
	tokens, err := ReadFile(filePath)
	if err != nil {
		log.Println("read the address list failed!")
	}
	tokenListLock.Lock()
	tokenList = *tokens
	tokenListLock.Unlock()
}