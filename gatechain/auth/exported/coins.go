package exported

// -----------------------------------------------------------------------------
// RevocableTx - record all Delay txs for account in special height
type RevocableTx struct {
	Txs    RevocableTxCoinsArray `json:"txs" yaml:"txs"`
	Height uint64                `json:"height" yaml:"height"`
}
