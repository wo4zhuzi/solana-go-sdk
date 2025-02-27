package rpc

import "encoding/json"

type Commitment string

const (
	CommitmentFinalized Commitment = "finalized"
	CommitmentConfirmed Commitment = "confirmed"
	CommitmentProcessed Commitment = "processed"
)

type ErrorResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type Context struct {
	Slot uint64 `json:"slot"`
}

type GeneralResponse struct {
	JsonRPC string         `json:"jsonrpc"`
	ID      uint64         `json:"id"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

type Instruction struct {
	ProgramIDIndex uint64   `json:"programIdIndex"`
	Accounts       []uint64 `json:"accounts"`
	Data           string   `json:"data"`
}

type InstructionJsonParsed struct {
	Parsed    json.RawMessage `json:"parsed"`
	Program   string          `json:"program"`
	ProgramId string          `json:"programId"`
	Accounts  *[]string       `json:"accounts"`
	Data      string          `json:"data"`
}

type Parsed struct {
	Info Info   `json:"info"`
	Type string `json:"type"`
}

type Info struct {
	Account     string      `json:"account"`
	NewAccount  string      `json:"newAccount"`
	Owner       string      `json:"owner"`
	Amount      string      `json:"amount"`
	Authority   string      `json:"authority"`
	Mint        string      `json:"mint"`
	Destination string      `json:"destination"`
	Lamports    float64     `json:"lamports"`
	Source      string      `json:"source"`
	TokenAmount TokenAmount `json:"tokenAmount"`
}

type TokenAmount struct {
	Amount         string  `json:"amount"`
	Decimals       int64   `json:"decimals"`
	UiAmount       float64 `json:"uiAmount"`
	UiAmountString string  `json:"uiAmountString"`
}

type TransactionMetaTokenBalance struct {
	AccountIndex  int    `json:"accountIndex"`
	Mint          string `json:"mint"`
	Owner         string `json:"owner"`
	ProgramId     string `json:"programId"`
	UITokenAmount struct {
		Amount         string `json:"amount"`
		Decimals       int64  `json:"decimals"`
		UIAmountString string `json:"uiAmountString"`
	} `json:"uiTokenAmount"`
}

type TransactionMeta struct {
	Fee               uint64                        `json:"fee"`
	PreBalances       []int64                       `json:"preBalances"`
	PostBalances      []int64                       `json:"postBalances"`
	PreTokenBalances  []TransactionMetaTokenBalance `json:"preTokenBalances"`
	PostTokenBalances []TransactionMetaTokenBalance `json:"postTokenBalances"`
	LogMessages       []string                      `json:"logMessages"`
	InnerInstructions []struct {
		Index        uint64        `json:"index"`
		Instructions []Instruction `json:"instructions"`
	} `json:"innerInstructions"`
	Err    interface{}            `json:"err"`
	Status map[string]interface{} `json:"status"`
}

type TransactionMetaJsonParsed struct {
	Fee               uint64                        `json:"fee"`
	PreBalances       []int64                       `json:"preBalances"`
	PostBalances      []int64                       `json:"postBalances"`
	PreTokenBalances  []TransactionMetaTokenBalance `json:"preTokenBalances"`
	PostTokenBalances []TransactionMetaTokenBalance `json:"postTokenBalances"`
	LogMessages       []string                      `json:"logMessages"`
	InnerInstructions []struct {
		Index        uint64                  `json:"index"`
		Instructions []InstructionJsonParsed `json:"instructions"`
	} `json:"innerInstructions"`
	Err    interface{}            `json:"err"`
	Status map[string]interface{} `json:"status"`
}

type MessageHeader struct {
	NumRequiredSignatures       uint8 `json:"numRequiredSignatures"`
	NumReadonlySignedAccounts   uint8 `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts uint8 `json:"numReadonlyUnsignedAccounts"`
}

type Message struct {
	Header          MessageHeader `json:"header"`
	AccountKeys     []string      `json:"accountKeys"`
	RecentBlockhash string        `json:"recentBlockhash"`
	Instructions    []Instruction `json:"instructions"`
}

type MessageJsonParsed struct {
	Header          MessageHeader           `json:"header"`
	AccountKeys     []AccountKeys           `json:"accountKeys"`
	RecentBlockhash string                  `json:"recentBlockhash"`
	Instructions    []InstructionJsonParsed `json:"instructions"`
}

type AccountKeys struct {
	Pubkey   string `json:"pubkey"`
	Signer   bool   `json:"signer"`
	Writable bool   `json:"writable"`
}

type Transaction struct {
	Signatures []string          `json:"signatures"`
	Message    MessageJsonParsed `json:"message"`
}

type Encoding string

const (
	EncodingBase58     Encoding = "base58" // limited to Account data of less than 128 bytes
	EncodingBase64     Encoding = "base64"
	EncodingBase64Zstd Encoding = "base64+zstd"
)
