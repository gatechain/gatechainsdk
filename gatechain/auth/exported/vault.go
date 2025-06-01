package exported

import (
	"fmt"
	
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

//-----------------------------------------------------------------------------
// RevocableTxCoins

// RevocableTxCoinsArray revocabledTxCoins Array
type RevocableTxCoinsArray []RevocableTxCoins

func (txs RevocableTxCoinsArray) String() string {
	result := fmt.Sprintln(`Txs: count `, len(txs))

	for _, data := range txs {
		result = result + fmt.Sprintf(`  TxHash:         REVOCABLEPAY-%s
  Index:          %d
  Height:         %d
  Tokens:          %s
`,

			data.TxHash, data.Index, data.Height, data.Coins,
		)
	}

	return result
}

// RevocableTxCoins - addRevocableTokens input struct
type RevocableTxCoins struct {
	TxHash string      `json:"tx_hash" yaml:"tx_hash"`
	Height uint64      `json:"height" yaml:"height"`
	Index  uint64      `json:"msg_index" yaml:"msg_index"`
	Coins  types.Coins `json:"tokens" yaml:"tokens"`
}

// VaultAccount defines an account type that can send innrecovable tx.
type VaultAccount interface {
	Account

	GetDelayHeight() uint64

	GetSecurityAddress() types.AccTypeAddress

	GetRevocableTokensDetail(height int64) (RevocableTxCoinsArray, error)

	//MergeRevocableTokens(height uint64) error
	MergeRevocableWei(height uint64) error
}
