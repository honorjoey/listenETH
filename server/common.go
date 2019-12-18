package server

import "sync"

var (
	//以太坊节点
	Net string

	//最新区块
	latestBlockNumber     int64
	latestBlockNumberLock sync.RWMutex

	//开始同步的区块号
	startBlockNumber int64

	//地址列表
	addressList             []string
	addressListLock         sync.RWMutex
	addressListFilePath     string
	addressListFilePathLock sync.RWMutex

	//token列表
	tokenList             []string
	tokenListLock         sync.RWMutex
	tokenListFilePath     string
	tokenListFilePathLock sync.RWMutex

	//邮箱
	EmailUser     string // 邮箱账号
	EmailPassword string //注意，此处为授权码、不是密码
	EmailHost     string //smtp地址及端口
	EmailTo       string //接收者，内容可重复，邮箱之间用；隔开
	EmailToLock   sync.RWMutex
)
