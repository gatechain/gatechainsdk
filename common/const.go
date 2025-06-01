package common

var (
	APIToken     = "your_api_token"
	EndPoint     = "http://your_server_ip:port"
	UnsignTxFile = "tx_unsign.json"
	SignTxFile   = "tx_sign.json"
)

const (
	// DefaultKeyPass contains the default key password for genesis transactions
	DefaultKeyPass = "12345678"
	MinKeyPassLen  = 8
)
