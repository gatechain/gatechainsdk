package types

import (
	"errors"
	"fmt"
)

// -----------------------------------------------------------------------------
// ClearingHeight
type ClearingHeight struct {
	LastClearingHeight       uint64 `json:"last_clearing_height"  yaml:"last_clearing_height"`
	LastClearingEffectHeight uint64 `json:"last_clearing_effect_height"  yaml:"last_clearing_effect_height"`
	NextClearingHeight       uint64 `json:"next_clearing_height"  yaml:"next_clearing_height"`               // the next clear time.If it is puls 0,user can not change clear time
	NextClearingEffectHeight uint64 `json:"next_clearing_effect_height"  yaml:"next_clearing_effect_height"` // the effective height of next clear time
}

// when ClearingHeight init,the nextClearingHeight and NextClearingEffectHeight must be 0.
func newClearingHeight(lastClearingHeight uint64, lastClearingEffectHeight uint64) *ClearingHeight {
	return &ClearingHeight{LastClearingHeight: lastClearingHeight, LastClearingEffectHeight: lastClearingEffectHeight, NextClearingHeight: 0, NextClearingEffectHeight: 0}
}

// get clearingHeight.
// If nextClearingHeight had effect,then return nextClearingHeight.
func (c *ClearingHeight) getClearingHeight(blockHeight uint64) uint64 {
	if c.isNextEffect(blockHeight) {
		//last setClearingHeight have effect
		return c.NextClearingHeight
	}
	return c.LastClearingHeight
}

// needClear  -  return this account whether clear time is reached
func (c *ClearingHeight) NeedClear(blockHeight uint64) bool {
	if c.isNextEffect(blockHeight) {
		//last setClearingHeight have effect
		return c.NextClearingHeight <= blockHeight
	}

	return (c.LastClearingHeight <= blockHeight) && (c.LastClearingHeight > 0)
}

// set clearingHeight.
// If nextClearingHeight had effect,then can set clearingHeight.
// If NextClearingEffectHeight == 0 (ClearingHeight had not be set clearingHeight),then then can set clearingHeight.
func (c *ClearingHeight) SetClearingHeight(clearingHeight uint64, effectHeight uint64, blockHeight uint64, delayArriveTime uint64) error {
	if delayArriveTime+blockHeight > clearingHeight {
		return errors.New("set clearing height failed, new clearingHeight must more than effectHeight of this  setClearingHeight ")
	}

	if c.IsCanSet(blockHeight, effectHeight) {
		c.LastClearingHeight = c.NextClearingHeight
		c.LastClearingEffectHeight = c.NextClearingEffectHeight
		c.NextClearingHeight = clearingHeight
		c.NextClearingEffectHeight = effectHeight
		return nil
	}

	//the account have ineffective clear time.
	return errors.New("set clear Time failed")
}

// Check nextClear have effected with current block height.
// If NextClearingEffectHeight > 0 (ClearingHeight had be set clearingHeight),then check block have reached to nextEffectHeight.
// If NextClearingEffectHeight == 0 (ClearingHeight had not be set clearingHeight),then nextClear.
func (c *ClearingHeight) isNextEffect(blockHeight uint64) bool {
	return c.NextClearingEffectHeight <= blockHeight && c.NextClearingEffectHeight > 0
}

// IsCanSet return whether this clear time can be set
func (c *ClearingHeight) IsCanSet(blockHeight uint64, effectHeight uint64) bool {
	if effectHeight < blockHeight {
		return false
	}

	return (c.isNextEffect(blockHeight)) || (c.NextClearingEffectHeight == 0)
}

func (c ClearingHeight) GetString() string {
	return fmt.Sprintf("last_clearing_height(%d) last_clearing_effect_height:(%d) next_clearing_height:(%d) next_clearing_effect_height:(%d)",
		c.LastClearingHeight, c.LastClearingEffectHeight, c.NextClearingHeight, c.NextClearingEffectHeight)
}
