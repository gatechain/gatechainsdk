package types

import (
	"fmt"
	"time"

	"github.com/gatechain/gatechainsdk/gatechain/types"
)

type (
	// Commission defines a commission parameters for a given validator.
	Commission struct {
		CommissionRates `json:"commission_rates" yaml:"commission_rates"`
		UpdateTime      time.Time `json:"update_time" yaml:"update_time"` // the last time the commission rate was changed
	}

	// CommissionRates defines the initial commission rates to be used for creating a
	// validator.
	CommissionRates struct {
		Rate          types.Dec `json:"rate" yaml:"rate"`                       // the commission rate charged to delegators, as a fraction
		MaxRate       types.Dec `json:"max_rate" yaml:"max_rate"`               // maximum commission rate which validator can ever charge, as a fraction
		MaxChangeRate types.Dec `json:"max_change_rate" yaml:"max_change_rate"` // maximum daily increase of the validator commission, as a fraction
	}
)

// String implements the Stringer interface for a Commission.
func (c Commission) String() string {
	return fmt.Sprintf("rate: %s, maxRate: %s, maxChangeRate: %s, updateTime: %s",
		c.Rate, c.MaxRate, c.MaxChangeRate, c.UpdateTime,
	)
}
