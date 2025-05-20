package types

// Transactions messages must fulfill the Msg
type Msg interface {

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Type() string

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []AccAddress
}

//__________________________________________________________

// Transactions objects must fulfill the Tx
type Tx interface {
	// Gets the all the transaction's messages.
	GetMsgs() []Msg
}

//__________________________________________________________

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)
