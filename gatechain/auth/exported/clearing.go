package exported

// -----------------------------------------------------------------------------
// ClearingHeight
type ClearingHeight struct {
	LastClearingHeight       uint64 `json:"last_clearing_height"  yaml:"last_clearing_height"`
	LastClearingEffectHeight uint64 `json:"last_clearing_effect_height"  yaml:"last_clearing_effect_height"`
	NextClearingHeight       uint64 `json:"next_clearing_height"  yaml:"next_clearing_height"`               // the next clear time.If it is puls 0,user can not change clear time
	NextClearingEffectHeight uint64 `json:"next_clearing_effect_height"  yaml:"next_clearing_effect_height"` // the effective height of next clear time
}
