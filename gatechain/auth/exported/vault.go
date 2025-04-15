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

	GetAccountType() uint8
	GetClear() ClearingHeight

	GetDelayHeight() uint64
	SetDelayHeight(uint64) error

	GetSecurityAddress() types.AccTypeAddress
	SetSecurityAddress(types.AccTypeAddress) error // errors if already set.

	GetVaultAddress() []types.AccTypeAddress
	AddVaultAddress(types.AccTypeAddress) error // errors if already set.

	GetRevocabledTokensTotal() types.Coins
	GetRevocabledTokens() []*RevocableTx
	GetRevocableTokensDetail(height int64) (RevocableTxCoinsArray, error)

	AddRevocabledTokens(RevocableTxCoins) error
	AddRevocableTokens(RevocableTxCoins) error
	ExistRevocabledTokens(RevocableTxCoins) bool
	ExistRevocableTokens(RevocableTxCoins) bool
	GetRevocableTokens() []*RevocableTx

	NeedClear(height uint64) bool
	IsCanSetClearingHeight(blockHeight uint64, effectHeight uint64) bool
	SetClearingHeight(clearingHeight uint64, effectHeight uint64, blockHeight uint64, delayArriveHeight uint64) error

	RemoveRevocabledTokens(RevocableTxCoins) error
	RemoveRevocableTokens(RevocableTxCoins) error

	MergeRevocableTokens(height uint64) error
	MergeRevocableWei(height uint64) error
}
