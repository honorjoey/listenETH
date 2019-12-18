package server

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/funkygao/golib/context"
	"github.com/joeylaker/loadConfig"
	"log"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

func TestSubstring(t *testing.T)  {
	s := "0x000000000000000000000000f88fa02d04c83fe6c5a4925d70025e2c76e68ca7"
	fmt.Println("0x"+s[26:])

	v := "0x000000000000000000000000000000000000000000000004d1753117f4b80000"
	val := strings.TrimPrefix(v, "0x")
	value, b := new(big.Int).SetString(val, 16)
	if !b {
		log.Println("getTxByHash json unmarshal failed!", b)
	}
	float := big.NewFloat(0).SetInt(value)
	fmt.Println(float)
	Wei, _ := new(big.Float).SetString("1000000000000000000")
	float = float.Quo(float, Wei)
	fmt.Println(float)
}

func TestAddress(t *testing.T) {
	//client, err := ethclient.Dial("https://mainnet.infura.io")
	client, err := ethclient.Dial("http://52.221.201.15:8232")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress("0x2b959ef258370c7a554d2bb052b3bc062d17e758")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address = common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}

func TestReadFile(t *testing.T) {
	config := loadConfig.NewConfig()
	config.LoadConfig("./config.ini")
	Net := config.GetConfig("ETHNode")
	start := config.GetConfig("startBlockNumber")
	number, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		log.Println("config blockNumber is not a number")
	}
	startBlockNumber := number
	addressListFilePath := config.GetConfig("addressListFile")
	tokenListFilePath := config.GetConfig("tokenListFile")
	//Email
	EmailUser := config.GetConfig("EmailUser")
	EmailPassword := config.GetConfig("EmailPassword")
	EmailHost := config.GetConfig("EmailHost")
	EmailTo := config.GetConfig("EmailTo")
	fmt.Println(Net)
	fmt.Println(startBlockNumber)
	fmt.Println(addressListFilePath)
	fmt.Println(tokenListFilePath)
	fmt.Println(EmailUser)
	fmt.Println(EmailPassword)
	fmt.Println(EmailHost)
	fmt.Println(EmailTo)
}

func TestReadAddressList(t *testing.T) {
	valueInt, err := strconv.ParseInt(strings.TrimPrefix("0x000000000000000000000000000000000000000000000004c53ecdc18a600000", "0x"), 16, 64)
	if err != nil {
		log.Println("hex token value to number failed!")
	}
	fmt.Println(valueInt)
}

func TestReceipt(t *testing.T) {
	requestStr := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionReceipt\",\"params\":[\"%s\"],\"id\":1}", "0xd752e6de92bd8e591245a9d19b676663c3f526df37073dffd3b3fa1f93ad45f0")
	respStr, err := request(requestStr, "http://47.92.223.116:8232", "POST")
	if err != nil {
		log.Println("getTxByHash.request", err)
	}
	resp := &TransactionReceipt{}
	err = json.Unmarshal([]byte(respStr), resp)
	if err != nil {
		log.Println("getTxByHash json unmarshal failed!", err)
	}
	fmt.Println("address", resp.Result.To)
}

func TestSendMail(t *testing.T) {
	subject := "以太坊账户变动通知"
	body := `
    <html>
    <body>
    <h3>
    "以太坊账户变动通知"
    </h3>
	<table>
		<tr>
			<td>From</td>
			<td>`+ "9999" +`</td>
		</tr>
		<tr>
			<td>To</td>
			<td>`+ "8888" +`</td>
		</tr>
		<tr>
			<td>Value</td>
			<td>`+ "5555" +`</td>
		</tr>
	</table>
    </body>
    </html>
    `
	err := SendMail("xxx@126.com", "xxx", "smtp.126.com:25", "xxx@126.com", subject, body, "html")
	if err != nil {
		fmt.Println("发送邮件失败!", EmailTo)
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功!", EmailTo)
	}
}

func TestTokenAddress(t *testing.T) {
	isContractAddress("")
}
