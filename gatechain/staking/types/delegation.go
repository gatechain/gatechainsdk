package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// Delegation represents the bond with tokens held by an account. It is
// owned by one delegator, and is associated with the voting power of one
// validator.
type Delegation struct {
	DelegatorAddress types.AccAddress `json:"delegator_address" yaml:"delegator_address"`
	ValidatorAddress types.ValAddress `json:"con-account_address" yaml:"con-account_address"`
	Shares           types.Dec        `json:"shares" yaml:"shares"`
}

// String returns a human readable string representation of a Delegation.
func (d Delegation) String() string {
	return fmt.Sprintf(`Delegation:
 Delegator:   %s
 Con-account: %s
 Shares:      %s`, d.DelegatorAddress,
		d.ValidatorAddress, d.Shares)
}

// UnbondingDelegation stores all of a single delegator's unbonding bonds
// for a single validator in an time-ordered list
type UnbondingDelegation struct {
	DelegatorAddress types.AccAddress           `json:"delegator_address" yaml:"delegator_address"`     // delegator
	ValidatorAddress types.ValAddress           `json:"con-account_address" yaml:"con-account_address"` // validator unbonding from operator addr
	Entries          []UnbondingDelegationEntry `json:"entries" yaml:"entries"`                         // unbonding delegation entries
}

// UnbondingDelegationEntry - entry to an UnbondingDelegation
type UnbondingDelegationEntry struct {
	CreationHeight int64     `json:"creation_height" yaml:"creation_height"` // height which the unbonding took place
	CompletionTime time.Time `json:"completion_time" yaml:"completion_time"` // time at which the unbonding delegation will complete
	InitialBalance types.Int `json:"initial_balance" yaml:"initial_balance"` // atoms initially scheduled to receive at completion
	Balance        types.Int `json:"balance" yaml:"balance"`                 // atoms to receive at completion
}

// String returns a human readable string representation of an UnbondingDelegation.
func (d UnbondingDelegation) String() string {
	out := fmt.Sprintf(`Undelegations  between:
  Delegator:   %s
  Con-account: %s
	Entries:`, d.DelegatorAddress, d.ValidatorAddress)
	for i, entry := range d.Entries {
		out += fmt.Sprintf(`    
		Undelegation         %d:
		Creation Height:           %v
		Min time to undelegate (unix): %v
		Expected balance:          %s`, i, entry.CreationHeight,
			entry.CompletionTime, entry.Balance)
	}
	return out
}

// UnbondingDelegations is a collection of UnbondingDelegation
type UnbondingDelegations []UnbondingDelegation

func (ubds UnbondingDelegations) String() (out string) {
	for _, u := range ubds {
		out += u.String() + "\n"
	}
	return strings.TrimSpace(out)
}

// Redelegation contains the list of a particular delegator's
// redelegating bonds from a particular source validator to a
// particular destination validator
type Redelegation struct {
	DelegatorAddress    types.AccAddress    `json:"delegator_address" yaml:"delegator_address"`             // delegator
	ValidatorSrcAddress types.ValAddress    `json:"con-account_src_address" yaml:"con-account_src_address"` // validator redelegation source operator addr
	ValidatorDstAddress types.ValAddress    `json:"con-account_dst_address" yaml:"con-account_dst_address"` // validator redelegation destination operator addr
	Entries             []RedelegationEntry `json:"entries" yaml:"entries"`                                 // redelegation entries
}

// RedelegationEntry - entry to a Redelegation
type RedelegationEntry struct {
	CreationHeight int64     `json:"creation_height" yaml:"creation_height"` // height at which the redelegation took place
	CompletionTime time.Time `json:"completion_time" yaml:"completion_time"` // time at which the redelegation will complete
	InitialBalance types.Int `json:"initial_balance" yaml:"initial_balance"` // initial balance when redelegation started
	SharesDst      types.Dec `json:"shares_dst" yaml:"shares_dst"`           // amount of destination-validator shares created by redelegation
}

// String returns a human readable string representation of a Redelegation.
func (d Redelegation) String() string {
	out := fmt.Sprintf(`Redelegations between:
  Delegator:                   %s
  Source Con-account:          %s
  Destination Con-account:     %s
  Entries:
`,
		d.DelegatorAddress, d.ValidatorSrcAddress, d.ValidatorDstAddress,
	)

	for i, entry := range d.Entries {
		out += fmt.Sprintf(`    Redelegation Entry #%d:
      Creation height:           %v
      Min time to undelegate (unix): %v
      Dest Shares:               %s
`,
			i, entry.CreationHeight, entry.CompletionTime, entry.SharesDst,
		)
	}

	return strings.TrimRight(out, "\n")
}

// ----------------------------------------------------------------------------
// client Types

// DelegationResponse is equivalent to Delegation except that it contains a balance
// in addition to shares which is more suitable for client responses.
type DelegationResponse struct {
	Delegation
	Balance types.Int `json:"balance" yaml:"balance"`
}

// String implements the Stringer interface for DelegationResponse.
func (d DelegationResponse) String() string {
	return fmt.Sprintf("%s\n  Balance:   %s", d.Delegation.String(), d.Balance)
}

type delegationRespAlias DelegationResponse

// MarshalJSON implements the json.Marshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (d DelegationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((delegationRespAlias)(d))
}

// UnmarshalJSON implements the json.Unmarshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (d *DelegationResponse) UnmarshalJSON(bz []byte) error {
	return json.Unmarshal(bz, (*delegationRespAlias)(d))
}

// DelegationResponses is a collection of DelegationResp
type DelegationResponses []DelegationResponse

// String implements the Stringer interface for DelegationResponses.
func (d DelegationResponses) String() (out string) {
	for _, del := range d {
		out += del.String() + "\n"
	}
	return strings.TrimSpace(out)
}

// RedelegationResponse is equivalent to a Redelegation except that its entries
// contain a balance in addition to shares which is more suitable for client
// responses.
type RedelegationResponse struct {
	Redelegation
	Entries []RedelegationEntryResponse `json:"entries" yaml:"entries"` // nolint: structtag
}

// RedelegationEntryResponse is equivalent to a RedelegationEntry except that it
// contains a balance in addition to shares which is more suitable for client
// responses.
type RedelegationEntryResponse struct {
	RedelegationEntry
	Balance types.Int `json:"balance"`
}

// String implements the Stringer interface for RedelegationResp.
func (r RedelegationResponse) String() string {
	out := fmt.Sprintf(`Redelegations between:
  Delegator:                   %s
  Source Con-account:          %s
  Destination Con-account:     %s
  Entries:
`,
		r.DelegatorAddress, r.ValidatorSrcAddress, r.ValidatorDstAddress,
	)

	for i, entry := range r.Entries {
		out += fmt.Sprintf(`    Redelegation Entry #%d:
      Creation height:           %v
      Min time to undelegate (unix): %v
      Initial Balance:           %s
      Shares:                    %s
      Balance:                   %s
`,
			i, entry.CreationHeight, entry.CompletionTime, entry.InitialBalance, entry.SharesDst, entry.Balance,
		)
	}

	return strings.TrimRight(out, "\n")
}

type redelegationRespAlias RedelegationResponse

// MarshalJSON implements the json.Marshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (r RedelegationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((redelegationRespAlias)(r))
}

// UnmarshalJSON implements the json.Unmarshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (r *RedelegationResponse) UnmarshalJSON(bz []byte) error {
	return json.Unmarshal(bz, (*redelegationRespAlias)(r))
}

// RedelegationResponses are a collection of RedelegationResp
type RedelegationResponses []RedelegationResponse

func (r RedelegationResponses) String() (out string) {
	for _, red := range r {
		out += red.String() + "\n"
	}
	return strings.TrimSpace(out)
}
