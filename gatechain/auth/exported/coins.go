package exported

import (
	"errors"
)

//-----------------------------------------------------------------------------
// RevocableTx - record all Delay txs for account in special height
type RevocableTx struct {
	Txs    RevocableTxCoinsArray `json:"txs" yaml:"txs"`
	Height uint64                `json:"height" yaml:"height"`
}

// newRevocableTokens generate new delay coins
func NewRevocableTokens(txCoins RevocableTxCoins) *RevocableTx {
	txs := RevocableTxCoinsArray{txCoins}
	return &RevocableTx{Txs: txs, Height: txCoins.Height}
}

// AddRevocableTokens add delay coins to delay tx
func (delayTx *RevocableTx) AddRevocableTokens(txCoins RevocableTxCoins) error {
	if delayTx.Height != txCoins.Height {
		return errors.New("tx height error")
	}

	exist := false
	for _, tx := range delayTx.Txs {
		if txCoins.TxHash == tx.TxHash {
			exist = true
			return errors.New("do not support multi revocable pay in one tx")
		}
	}

	if !exist {
		index := 0
		breakOut := false
		tx := RevocableTxCoins{}
		for index, tx = range delayTx.Txs {
			if tx.TxHash > txCoins.TxHash {
				breakOut = true
				break
			}
		}
		if breakOut {
			rear := append([]RevocableTxCoins{}, delayTx.Txs[index:]...)
			delayTx.Txs = append(append(delayTx.Txs[:index], txCoins), rear...)
		} else {
			delayTx.Txs = append(delayTx.Txs, txCoins)
		}
	}

	return nil
}
