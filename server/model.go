package server

type (
	TransactionReceipt struct {
		JsonRPC string      `json:"jsonrpc"`
		Id      int64       `json:"id"`
		Result  ReceiptData `json:"result"`
	}

	ReceiptData struct {
		BlockHash         string   `json:"blockHash"`
		BlockNumber       string   `json:"blockNumber"`
		ContractAddress   string   `json:"contractAddress"`
		CumulativeGasUsed string   `json:"cumulativeGasUsed"`
		From              string   `json:"from"`
		GasUsed           string   `json:"gasUsed"`
		Logs              []TxLogs `json:"logs"`
		LogsBloom         string   `json:"logsBloom"`
		Status            string   `json:"status"`
		To                string   `json:"to"`
		TxHash            string   `json:"transactionHash"`
		TxIndex           string   `json:"transactionIndex"`
	}

	TxLogs struct {
		Address     string   `json:"address"`
		Topics      []string `json:"topics"`
		Data        string   `json:"data"`
		BlockNumber string   `json:"blockNumber"`
		TxHash      string   `json:"transactionHash"`
		TxIndex     string   `json:"transactionIndex"`
		BlockHash   string   `json:"blockHash"`
		LogIndex    string   `json:"logIndex"`
		Removed     bool     `json:"removed"`
	}

	Transaction struct {
		JsonRPC string      `json:"jsonrpc"`
		Id      int64       `json:"id"`
		Result  ETHTxResult `json:"result"`
	}

	ETHTxResult struct {
		Difficulty      string            `json:"difficulty"`
		ExtraData       string            `json:"extraData"`
		GasLimit        string            `json:"gasLimit"`
		GasUsed         string            `json:"gasUsed"`
		Hash            string            `json:"hash"`
		LogsBloom       string            `json:"logsBloom"`
		Miner           string            `json:"miner"`
		MixHash         string            `json:"mixHash"`
		Nonce           string            `json:"nonce"`
		Number          string            `json:"number"`
		ParentHash      string            `json:"parentHash"`
		ReceiptsRoot    string            `json:"receiptsRoot"`
		Sha3Uncles      string            `json:"sha3Uncles"`
		Size            string            `json:"size"`
		StateRoot       string            `json:"stateRoot"`
		Timestamp       string            `json:"timestamp"`
		TotalDifficulty string            `json:"totalDifficulty"`
		Txs             []TransactionData `json:"transactions"`
		TxRoot          string            `json:"transactionsRoot"`
		Uncles          []string          `json:"uncles"`
	}

	BlockNumberResp struct {
		JsonRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  string `json:"result"`
	}

	TransactionData struct {
		Time  int64  `json:"time"`
		From  string `json:"from"`
		To    string `json:"to"`
		Value string `json:"value"`
		Hash  string `json:"hash"`
	}

	ETHEmail struct {
		From         string
		To           string
		Value        string
		TxType       string
		TokenAddress string
	}
)
