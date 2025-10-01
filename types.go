package ethrpc

import (
	"bytes"
	"encoding/json"
	"math/big"
)

// Syncing - object with syncing data info
type Syncing struct {
	IsSyncing     bool
	StartingBlock int
	CurrentBlock  int
	HighestBlock  int
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Syncing) UnmarshalJSON(data []byte) error {
	proxy := new(proxySyncing)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	// 手动赋值，避免使用 unsafe.Pointer
	s.IsSyncing = true
	s.StartingBlock = int(proxy.StartingBlock)
	s.CurrentBlock = int(proxy.CurrentBlock)
	s.HighestBlock = int(proxy.HighestBlock)

	return nil
}

// T - input transaction object
type T struct {
	From     string
	To       string
	Gas      int
	GasPrice *big.Int
	Value    *big.Int
	Data     string
	Nonce    int
}

// MarshalJSON implements the json.Unmarshaler interface.
func (t T) MarshalJSON() ([]byte, error) {
	params := map[string]interface{}{}
	if t.To != "" {
		params["to"] = t.To
	}
	if t.From != "" {
		params["from"] = t.From
	}
	if t.Gas > 0 {
		params["gas"] = IntToHex(t.Gas)
	}
	if t.GasPrice != nil {
		params["gasPrice"] = BigToHex(*t.GasPrice)
	}
	if t.Value != nil {
		params["value"] = BigToHex(*t.Value)
	}
	if t.Data != "" {
		params["data"] = t.Data
	}
	if t.Nonce > 0 {
		params["nonce"] = IntToHex(t.Nonce)
	}

	return json.Marshal(params)
}

// Transaction - transaction object
type Transaction struct {
	Hash             string
	Nonce            int
	BlockHash        string
	BlockNumber      *int
	TransactionIndex *int
	From             string
	To               string
	Value            big.Int
	Gas              int
	GasPrice         big.Int
	Input            string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransaction)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	// 手动赋值，避免使用 unsafe.Pointer
	t.Hash = proxy.Hash
	t.Nonce = int(proxy.Nonce)
	t.BlockHash = proxy.BlockHash
	if proxy.BlockNumber != nil {
		blockNum := int(*proxy.BlockNumber)
		t.BlockNumber = &blockNum
	}
	if proxy.TransactionIndex != nil {
		txIndex := int(*proxy.TransactionIndex)
		t.TransactionIndex = &txIndex
	}
	t.From = proxy.From
	t.To = proxy.To
	t.Value = big.Int(proxy.Value)
	t.Gas = int(proxy.Gas)
	t.GasPrice = big.Int(proxy.GasPrice)
	t.Input = proxy.Input

	return nil
}

// Log - log object
type Log struct {
	Removed          bool
	LogIndex         int
	TransactionIndex int
	TransactionHash  string
	BlockNumber      int
	BlockHash        string
	Address          string
	Data             string
	Topics           []string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (log *Log) UnmarshalJSON(data []byte) error {
	proxy := new(proxyLog)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	// 手动赋值，避免使用 unsafe.Pointer
	log.Removed = proxy.Removed
	log.LogIndex = int(proxy.LogIndex)
	log.TransactionIndex = int(proxy.TransactionIndex)
	log.TransactionHash = proxy.TransactionHash
	log.BlockNumber = int(proxy.BlockNumber)
	log.BlockHash = proxy.BlockHash
	log.Address = proxy.Address
	log.Data = proxy.Data
	log.Topics = proxy.Topics

	return nil
}

// FilterParams - Filter parameters object
type FilterParams struct {
	FromBlock string     `json:"fromBlock,omitempty"`
	ToBlock   string     `json:"toBlock,omitempty"`
	Address   []string   `json:"address,omitempty"`
	Topics    [][]string `json:"topics,omitempty"`
}

// TransactionReceipt - transaction receipt object
type TransactionReceipt struct {
	TransactionHash   string
	TransactionIndex  int
	BlockHash         string
	BlockNumber       int
	CumulativeGasUsed int
	GasUsed           int
	ContractAddress   string
	Logs              []Log
	LogsBloom         string
	Root              string
	Status            string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TransactionReceipt) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransactionReceipt)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	// 手动赋值，避免使用 unsafe.Pointer
	t.TransactionHash = proxy.TransactionHash
	t.TransactionIndex = int(proxy.TransactionIndex)
	t.BlockHash = proxy.BlockHash
	t.BlockNumber = int(proxy.BlockNumber)
	t.CumulativeGasUsed = int(proxy.CumulativeGasUsed)
	t.GasUsed = int(proxy.GasUsed)
	t.ContractAddress = proxy.ContractAddress
	t.Logs = proxy.Logs
	t.LogsBloom = proxy.LogsBloom
	t.Root = proxy.Root
	t.Status = proxy.Status

	return nil
}

// Block - block object
type Block struct {
	Number           int
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	Miner            string
	Difficulty       big.Int
	TotalDifficulty  big.Int
	ExtraData        string
	Size             int
	GasLimit         int
	GasUsed          int
	Timestamp        int
	Uncles           []string
	Transactions     []Transaction
}

type proxySyncing struct {
	IsSyncing     bool   `json:"-"`
	StartingBlock hexInt `json:"startingBlock"`
	CurrentBlock  hexInt `json:"currentBlock"`
	HighestBlock  hexInt `json:"highestBlock"`
}

type proxyTransaction struct {
	Hash             string  `json:"hash"`
	Nonce            hexInt  `json:"nonce"`
	BlockHash        string  `json:"blockHash"`
	BlockNumber      *hexInt `json:"blockNumber"`
	TransactionIndex *hexInt `json:"transactionIndex"`
	From             string  `json:"from"`
	To               string  `json:"to"`
	Value            hexBig  `json:"value"`
	Gas              hexInt  `json:"gas"`
	GasPrice         hexBig  `json:"gasPrice"`
	Input            string  `json:"input"`
}

type proxyLog struct {
	Removed          bool     `json:"removed"`
	LogIndex         hexInt   `json:"logIndex"`
	TransactionIndex hexInt   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockNumber      hexInt   `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

type proxyTransactionReceipt struct {
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  hexInt `json:"transactionIndex"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       hexInt `json:"blockNumber"`
	CumulativeGasUsed hexInt `json:"cumulativeGasUsed"`
	GasUsed           hexInt `json:"gasUsed"`
	ContractAddress   string `json:"contractAddress,omitempty"`
	Logs              []Log  `json:"logs"`
	LogsBloom         string `json:"logsBloom"`
	Root              string `json:"root"`
	Status            string `json:"status,omitempty"`
}

type hexInt int

func (i *hexInt) UnmarshalJSON(data []byte) error {
	result, err := ParseInt(string(bytes.Trim(data, `"`)))
	*i = hexInt(result)

	return err
}

type hexBig big.Int

func (i *hexBig) UnmarshalJSON(data []byte) error {
	result, err := ParseBigInt(string(bytes.Trim(data, `"`)))
	*i = hexBig(result)

	return err
}

type proxyBlock interface {
	toBlock() Block
}

type proxyBlockWithTransactions struct {
	Number           hexInt             `json:"number"`
	Hash             string             `json:"hash"`
	ParentHash       string             `json:"parentHash"`
	Nonce            string             `json:"nonce"`
	Sha3Uncles       string             `json:"sha3Uncles"`
	LogsBloom        string             `json:"logsBloom"`
	TransactionsRoot string             `json:"transactionsRoot"`
	StateRoot        string             `json:"stateRoot"`
	Miner            string             `json:"miner"`
	Difficulty       hexBig             `json:"difficulty"`
	TotalDifficulty  hexBig             `json:"totalDifficulty"`
	ExtraData        string             `json:"extraData"`
	Size             hexInt             `json:"size"`
	GasLimit         hexInt             `json:"gasLimit"`
	GasUsed          hexInt             `json:"gasUsed"`
	Timestamp        hexInt             `json:"timestamp"`
	Uncles           []string           `json:"uncles"`
	Transactions     []proxyTransaction `json:"transactions"`
}

func (proxy *proxyBlockWithTransactions) toBlock() Block {
	// 手动赋值，避免使用 unsafe.Pointer
	block := Block{}
	block.Number = int(proxy.Number)
	block.Hash = proxy.Hash
	block.ParentHash = proxy.ParentHash
	block.Nonce = proxy.Nonce
	block.Sha3Uncles = proxy.Sha3Uncles
	block.LogsBloom = proxy.LogsBloom
	block.TransactionsRoot = proxy.TransactionsRoot
	block.StateRoot = proxy.StateRoot
	block.Miner = proxy.Miner
	block.Difficulty = big.Int(proxy.Difficulty)
	block.TotalDifficulty = big.Int(proxy.TotalDifficulty)
	block.ExtraData = proxy.ExtraData
	block.Size = int(proxy.Size)
	block.GasLimit = int(proxy.GasLimit)
	block.GasUsed = int(proxy.GasUsed)
	block.Timestamp = int(proxy.Timestamp)
	block.Uncles = proxy.Uncles

	// 转换 transactions
	block.Transactions = make([]Transaction, len(proxy.Transactions))
	for i, tx := range proxy.Transactions {
		block.Transactions[i] = Transaction{}
		block.Transactions[i].Hash = tx.Hash
		block.Transactions[i].Nonce = int(tx.Nonce)
		block.Transactions[i].BlockHash = tx.BlockHash
		if tx.BlockNumber != nil {
			blockNum := int(*tx.BlockNumber)
			block.Transactions[i].BlockNumber = &blockNum
		}
		if tx.TransactionIndex != nil {
			txIndex := int(*tx.TransactionIndex)
			block.Transactions[i].TransactionIndex = &txIndex
		}
		block.Transactions[i].From = tx.From
		block.Transactions[i].To = tx.To
		block.Transactions[i].Value = big.Int(tx.Value)
		block.Transactions[i].Gas = int(tx.Gas)
		block.Transactions[i].GasPrice = big.Int(tx.GasPrice)
		block.Transactions[i].Input = tx.Input
	}

	return block
}

type proxyBlockWithoutTransactions struct {
	Number           hexInt   `json:"number"`
	Hash             string   `json:"hash"`
	ParentHash       string   `json:"parentHash"`
	Nonce            string   `json:"nonce"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	LogsBloom        string   `json:"logsBloom"`
	TransactionsRoot string   `json:"transactionsRoot"`
	StateRoot        string   `json:"stateRoot"`
	Miner            string   `json:"miner"`
	Difficulty       hexBig   `json:"difficulty"`
	TotalDifficulty  hexBig   `json:"totalDifficulty"`
	ExtraData        string   `json:"extraData"`
	Size             hexInt   `json:"size"`
	GasLimit         hexInt   `json:"gasLimit"`
	GasUsed          hexInt   `json:"gasUsed"`
	Timestamp        hexInt   `json:"timestamp"`
	Uncles           []string `json:"uncles"`
	Transactions     []string `json:"transactions"`
}

func (proxy *proxyBlockWithoutTransactions) toBlock() Block {
	// 手动赋值，避免使用 unsafe.Pointer
	block := Block{}
	block.Number = int(proxy.Number)
	block.Hash = proxy.Hash
	block.ParentHash = proxy.ParentHash
	block.Nonce = proxy.Nonce
	block.Sha3Uncles = proxy.Sha3Uncles
	block.LogsBloom = proxy.LogsBloom
	block.TransactionsRoot = proxy.TransactionsRoot
	block.StateRoot = proxy.StateRoot
	block.Miner = proxy.Miner
	block.Difficulty = big.Int(proxy.Difficulty)
	block.TotalDifficulty = big.Int(proxy.TotalDifficulty)
	block.ExtraData = proxy.ExtraData
	block.Size = int(proxy.Size)
	block.GasLimit = int(proxy.GasLimit)
	block.GasUsed = int(proxy.GasUsed)
	block.Timestamp = int(proxy.Timestamp)
	block.Uncles = proxy.Uncles

	block.Transactions = make([]Transaction, len(proxy.Transactions))
	for i := range proxy.Transactions {
		block.Transactions[i] = Transaction{}
		block.Transactions[i].Hash = proxy.Transactions[i]
	}

	return block
}
