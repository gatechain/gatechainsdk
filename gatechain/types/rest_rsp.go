package types

type CalculateGasResponse struct {
	Estimate uint64 `json:"estimate,omitempty"`
	Adjusted uint64 `json:"adjusted,omitempty"`
}

type RspAccount struct {
	Name          string    `json:"name" yaml:"name"`
	Coins         []RspCoin `json:"tokens" yaml:"tokens"`
	Address       string    `json:"address" yaml:"address"`
	AccountNumber uint64    `json:"account_number" yaml:"account_number"`
	Sequence      uint64    `json:"sequence" yaml:"sequence"`
}

type RspCoin struct {
	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3,customtype=Int" json:"amount"`
}

type RspStatus struct {
	ChainId   string `json:"chainId"`
	LastRound uint64 `json:"lastHeight"`
}

type RspBroadcastTx struct {
	TxHash string `json:"txhash,omitempty"`
	Code   uint32 `json:"code,omitempty"`
	Data   string `json:"data,omitempty"`
	Height uint64 `json:"height,omitempty"`
	RawLog string `json:"raw_log,omitempty"`
}

type RspQueryTx struct {
	TxHash    string          `json:"txhash"`
	Code      uint32          `json:"code,omitempty"`
	Height    uint64          `json:"height,omitempty"`
	GasWanted uint64          `json:"gas_wanted,omitempty"`
	GasUsed   uint64          `json:"gas_used,omitempty"`
	Codespace string          `json:"codespace,omitempty"`
	Logs      ABCIMessageLogs `json:"logs,omitempty"`
	RawLog    string          `json:"raw_log,omitempty"`
}
