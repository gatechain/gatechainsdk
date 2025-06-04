package client

import (
	"fmt"
	"strings"

	"github.com/gatechain/gatechainsdk/gatechain/rpc/spec/v1"
	"github.com/gatechain/gatechainsdk/gatechain/staking/types"
	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

const (
	OnLine  = "online"
	OffLine = "offline"
)

type TmpValidator struct {
	OperatorAddress sdk.ValAddress    `json:"operator_address" yaml:"operator_address"` // the bech32 address of the validator's operator
	Power           uint64            `json:"power" yaml:"power"`
	Status          string            `json:"status" yaml:"status"`
	PubKey          string            `json:"pubkey" yaml:"pubkey"`                     // the bech32 public key of the validator
	Tokens          sdk.Int           `json:"tokens" yaml:"tokens"`                     // delegated tokens
	PowerRate       sdk.Dec           `json:"power_rate" yaml:"power_rate"`             // poewe rate
	DelegatorShares sdk.Dec           `json:"delegator_shares" yaml:"delegator_shares"` // total shares issued to a validator's delegators
	Description     types.Description `json:"description" yaml:"description"`           // description terms for the validator
	Commission      types.Commission  `json:"commission" yaml:"commission"`             // commission parameters
}

func GetTmpValidator(v types.Validator, gmValidator v1.ResultConAccount) (*TmpValidator, error) {
	bechConsPubKey, err := sdk.Bech32ifyConsPub(v.ConsPubKey)
	if err != nil {
		return nil, fmt.Errorf("validator(%s) public key error:%s", v.OperatorAddress.String(), v.ConsPubKey)
	}
	status := OffLine
	if gmValidator.IsFind() {
		status = OnLine
	} else {
		v.PowerRate = sdk.NewDec(1)
	}
	return &TmpValidator{
		Power:           gmValidator.Power,
		OperatorAddress: v.OperatorAddress,
		Status:          status,
		PubKey:          bechConsPubKey,
		Tokens:          v.Tokens,
		DelegatorShares: v.DelegatorShares,
		PowerRate:       v.PowerRate,
		Description:     v.Description,
		Commission:      v.Commission,
	}, nil
}

// Validators is a collection of Validator
type TmpValidators []TmpValidator

func (v TmpValidators) String() (out string) {
	for _, val := range v {
		out += val.String() + "\n"
	}
	return strings.TrimSpace(out)
}

// String returns a human readable string representation of a validator.
func (v TmpValidator) String() string {
	return fmt.Sprintf(`Con-account
  Address:                    %s
  Pubkey:                     %s
  Status:                     %s
  Power:                      %v
  Tokens:                     %s
  Delegator Shares:           %s
  Power Rate:                 %s
  Description:                %s
  Commission:                 %s`, v.OperatorAddress, v.PubKey, v.Status,
		v.Power, v.Tokens,
		v.DelegatorShares, v.PowerRate, v.Description, v.Commission)
}

type TmpLocalKey struct {
}
