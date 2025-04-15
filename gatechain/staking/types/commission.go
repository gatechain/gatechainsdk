package types

import (
	"fmt"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
	"time"
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
		Rate          types2.Dec `json:"rate" yaml:"rate"`                       // the commission rate charged to delegators, as a fraction
		MaxRate       types2.Dec `json:"max_rate" yaml:"max_rate"`               // maximum commission rate which validator can ever charge, as a fraction
		MaxChangeRate types2.Dec `json:"max_change_rate" yaml:"max_change_rate"` // maximum daily increase of the validator commission, as a fraction
	}
)

// NewCommissionRates returns an initialized validator commission rates.
func NewCommissionRates(rate, maxRate, maxChangeRate types2.Dec) CommissionRates {
	return CommissionRates{
		Rate:          rate,
		MaxRate:       maxRate,
		MaxChangeRate: maxChangeRate,
	}
}

// NewCommission returns an initialized validator commission.
func NewCommission(rate, maxRate, maxChangeRate types2.Dec) Commission {
	return Commission{
		CommissionRates: NewCommissionRates(rate, maxRate, maxChangeRate),
		UpdateTime:      time.Unix(0, 0).UTC(),
	}
}

// NewCommission returns an initialized validator commission with a specified
// update time which should be the current block BFT time.
func NewCommissionWithTime(rate, maxRate, maxChangeRate types2.Dec, updatedAt time.Time) Commission {
	return Commission{
		CommissionRates: NewCommissionRates(rate, maxRate, maxChangeRate),
		UpdateTime:      updatedAt,
	}
}

// Equal checks if the given Commission object is equal to the receiving
// Commission object.
func (c Commission) Equal(c2 Commission) bool {
	return c.Rate.Equal(c2.Rate) &&
		c.MaxRate.Equal(c2.MaxRate) &&
		c.MaxChangeRate.Equal(c2.MaxChangeRate) &&
		c.UpdateTime.Equal(c2.UpdateTime)
}

// String implements the Stringer interface for a Commission.
func (c Commission) String() string {
	return fmt.Sprintf("rate: %s, maxRate: %s, maxChangeRate: %s, updateTime: %s",
		c.Rate, c.MaxRate, c.MaxChangeRate, c.UpdateTime,
	)
}

// Validate performs basic sanity validation checks of initial commission
// parameters. If validation fails, an SDK error is returned.
func (c CommissionRates) Validate() types2.Error {
	switch {
	case c.MaxRate.LT(types2.ZeroDec()):
		// max rate cannot be negative
		return ErrCommissionNegative(DefaultCodespace)

	case c.MaxRate.GT(types2.OneDec()):
		// max rate cannot be greater than 1
		return ErrCommissionHuge(DefaultCodespace)

	case c.Rate.LT(types2.ZeroDec()):
		// rate cannot be negative
		return ErrCommissionNegative(DefaultCodespace)

	case c.Rate.GT(c.MaxRate):
		// rate cannot be greater than the max rate
		return ErrCommissionGTMaxRate(DefaultCodespace)

	case c.MaxChangeRate.LT(types2.ZeroDec()):
		// change rate cannot be negative
		return ErrCommissionChangeRateNegative(DefaultCodespace)

	case c.MaxChangeRate.GT(c.MaxRate):
		// change rate cannot be greater than the max rate
		return ErrCommissionChangeRateGTMaxRate(DefaultCodespace)
	}

	return nil
}

// ValidateNewRate performs basic sanity validation checks of a new commission
// rate. If validation fails, an SDK error is returned.
func (c Commission) ValidateNewRate(newRate types2.Dec, blockTime time.Time) types2.Error {
	switch {
	case blockTime.Sub(c.UpdateTime).Hours() < 24:
		// new rate cannot be changed more than once within 24 hours
		return ErrCommissionUpdateTime(DefaultCodespace)

	case newRate.LT(types2.ZeroDec()):
		// new rate cannot be negative
		return ErrCommissionNegative(DefaultCodespace)

	case newRate.GT(c.MaxRate):
		// new rate cannot be greater than the max rate
		return ErrCommissionGTMaxRate(DefaultCodespace)

	case newRate.Sub(c.Rate).GT(c.MaxChangeRate):
		// new rate % points change cannot be greater than the max change rate
		return ErrCommissionGTMaxChangeRate(DefaultCodespace)
	}

	return nil
}

func (c Commission) ValidateNewMaxRate(newMaxRate types2.Dec, newMaxChangeRate types2.Dec) types2.Error {
	switch {
	case newMaxRate.LT(types2.ZeroDec()):
		// new rate cannot be negative
		return ErrCommissionNegative(DefaultCodespace)
	case newMaxRate.LT(c.Rate):
		return ErrCommissionGTMaxRate(DefaultCodespace)
	case newMaxRate.GT(types2.OneDec()):
		// max rate cannot be greater than 1
		return ErrCommissionHuge(DefaultCodespace)
	case newMaxChangeRate.LTE(types2.ZeroDec()):
		return ErrCommissionChangeRateNegative(DefaultCodespace)
	case newMaxChangeRate.GT(newMaxRate):
		// change rate cannot be greater than the max rate
		return ErrCommissionChangeRateGTMaxRate(DefaultCodespace)
	}
	return nil
}
