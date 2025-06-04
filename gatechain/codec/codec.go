package codec

import (
	cryptoamino "github.com/gatechain/crypto/encoding/amino"
	"github.com/tendermint/go-amino"
)

// amino codec to marshal/unmarshal
type Codec = amino.Codec

func New() *Codec {
	return amino.NewCodec()
}

// Register the go-crypto to the codec
func RegisterCrypto(cdc *Codec) {
	cryptoamino.RegisterAmino(cdc)
}

//__________________________________________________________________

// generic sealed codec to be used throughout framework
var Cdc *Codec

func init() {
	cdc := New()
	RegisterCrypto(cdc)
	Cdc = cdc.Seal()
}
