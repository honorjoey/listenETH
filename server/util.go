package server

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/funkygao/golib/context"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func request(jsonString, url, model string) (string, error) {
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}

	req.Header.SetMethod(model)

	req.Header.SetContentType("application/json")
	req.SetBodyString(jsonString)

	req.SetRequestURI(url)

	err := fasthttp.Do(req, resp)
	if err != nil {
		log.Println("fasthttp.Do() ", err.Error())
		return "", err
	}
	return string(resp.Body()), nil
}

func getTxReceiptByHash(hash string) (*TransactionReceipt, error) {
	requestStr := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionReceipt\",\"params\":[\"%s\"],\"id\":1}", hash)
	respStr, err := request(requestStr, Net, "POST")
	if err != nil {
		log.Println("getTxByHash.request", err)
		return nil, err
	}
	resp := &TransactionReceipt{}
	err = json.Unmarshal([]byte(respStr), resp)
	if err != nil {
		log.Println("getTxByHash json unmarshal failed!", err)
		return nil, err
	}
	return resp, nil
}

func getTxByHash(hash string) (*TransactionData, error) {
	requestStr := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionByHash\",\"params\":[\"%s\"],\"id\":1}", hash)
	respStr, err := request(requestStr, Net, "POST")
	if err != nil {
		log.Println("getTxByHash.request", err)
		return nil, err
	}
	resp := &TransactionData{}
	err = json.Unmarshal([]byte(respStr), resp)
	if err != nil {
		log.Println("getTxByHash json unmarshal failed!", err)
		return nil, err
	}
	return resp, nil
}

func getBlockByNumber(blockNumber int64) (*[]TransactionData, error) {
	blockNum := "0x" + strconv.FormatInt(blockNumber, 16)
	requestStr := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"%s\", true],\"id\":1}", blockNum)
	respStr, err := request(requestStr, Net, "POST")
	if err != nil {
		log.Println("getBlock.request", err)
		return nil, err
	}
	resp := &Transaction{}
	err = json.Unmarshal([]byte(respStr), resp)
	if err != nil {
		log.Println("getBlock json unmarshal failed!", err)
		return nil, err
	}
	txs := resp.Result.Txs
	return &txs, nil
}

func getLatestBlockNumber()  {
	for {
		time.Sleep(500*time.Millisecond)
		res, err := request(fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\",\"params\":[],\"id\":1}"), Net, "POST")
		if err != nil {
			fmt.Println(err)
		}
		resp := BlockNumberResp{}
		err = json.Unmarshal([]byte(res), &resp)
		if err != nil {
			log.Println("json.Unmarshal()", err.Error())
		}
		num, err := strconv.ParseInt(strings.TrimPrefix(resp.Result, "0x"), 16, 64)
		if err != nil {
			fmt.Println("result is not a number!")
		}
		latestBlockNumberLock.Lock()
		latestBlockNumber = num
		latestBlockNumberLock.Unlock()
	}
}

func ReadFile(filename string) (*[]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("open the file failed!")
		return nil, err
	}
	defer file.Close()
	buf := make([]byte, 4096*1000)
	_, err = file.Read(buf)
	if err != nil {
		log.Println("read the file failed!")
		return nil, err
	}
	var bytes []byte
	for i, v := range buf {
		if v == 0 {
			bytes = buf[:i]
			break
		}
	}
	slice := strings.Split(string(bytes), "\n")

	return &slice, nil
}

func isContractAddress(address string) bool {
	client, err := ethclient.Dial("http://52.221.201.15:8232")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	addr := common.HexToAddress(address)
	bytecode, err := client.CodeAt(context.Background(), addr, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	if len(bytecode) > 0 {
		return true
	}
	return false
}
