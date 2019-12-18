package server

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"
)

var EmailSlice []ETHEmail

var d time.Duration

func getTxs()  {
	for {
		time.Sleep(50*time.Millisecond)
		fmt.Println(startBlockNumber)
		latestBlockNumberLock.RLock()
		latest := latestBlockNumber
		latestBlockNumberLock.RUnlock()
		if startBlockNumber  >= latest - 12 {
			continue
		}
		txs, err := getBlockByNumber(startBlockNumber)
		if err != nil {
			log.Println("getTxs failed", err)
		}
		for _, resp := range *txs {
			//如果不是合约地址
			if !isContractAddress(resp.To) {
				//同步以太币转账
				addressListLock.RLock()
				addressListBak := addressList
				addressListLock.RUnlock()
				for _, v := range addressListBak {
					if strings.ToLower(v) == strings.ToLower(resp.To) {
						value, b := new(big.Int).SetString(strings.TrimPrefix(resp.Value, "0x"), 16)
						if !b {
							log.Println("getTxByHash json unmarshal failed!", b)
						}
						float := big.NewFloat(0).SetInt(value)
						Wei, _ := new(big.Float).SetString("1000000000000000000")
						float = float.Quo(float, Wei)
						fmt.Println("[ETH]", resp.From, v, float)
						eth := ETHEmail{
							TxType:"ETH",
							TokenAddress:"",
							From:resp.From,
							To:v,
							Value:float.String(),
						}
						EmailSlice = append(EmailSlice, eth)
					}
				}
			}else{
				tokenListLock.RLock()
				tokenListBak := tokenList
				tokenListLock.RUnlock()
				for _, token := range tokenListBak {
					if strings.ToLower(resp.To) == strings.ToLower(token) {
						//同步代币转账
						receipt, err := getTxReceiptByHash(resp.Hash)
						if err != nil {
							log.Println("get receipt failed", err)
						}
						if len(receipt.Result.Logs) != 0 {
							for _, v := range receipt.Result.Logs {
								to := strings.ToLower("0x" + v.Topics[2][26:])
								addressListLock.RLock()
								addressListBak := addressList
								addressListLock.RUnlock()
								for _, addr := range addressListBak {
									if strings.ToLower(addr) == to {
										from := strings.ToLower("0x" + v.Topics[1][26:])
										value, b := new(big.Int).SetString(strings.TrimPrefix(v.Data, "0x"), 16)
										if !b {
											log.Println("getTxByHash json unmarshal failed!", b)
										}
										float := big.NewFloat(0).SetInt(value)
										Wei, _ := new(big.Float).SetString("1000000000000000000")
										float = float.Quo(float, Wei)
										eth := ETHEmail{
											TxType:"Token",
											TokenAddress:token,
											From:from,
											To:to,
											Value:float.String(),
										}
										EmailSlice = append(EmailSlice, eth)
									}
								}
							}
						}
					}
				}
			}
		}
		startBlockNumber ++
	}
}

func sendEThMessage() {
	d, _ = time.ParseDuration("60s")
	t := time.NewTicker(d)
	for {
		<-t.C
		if len(EmailSlice) > 0 {
			sendTxs(&EmailSlice)
			EmailSlice = EmailSlice[0:0]
		}else{
			continue
		}
	}
}
