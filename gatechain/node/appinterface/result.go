package appinterface

import "fmt"

const (
	CodeTypeOK            uint32 = 0
	CodeTypeError         uint32 = 1
	CodeTypeParticipation uint32 = 2
)

func (r ResponseStatus) IsOK() bool {
	return r.Code == CodeTypeOK
}

func (r ResponseStatus) IsErr() bool {
	return r.Code != CodeTypeOK && r.Code != CodeTypeParticipation
}

func (r ResponseStatus) IsParticipation() bool {
	return r.Code == CodeTypeParticipation
}

// IsOK returns true if Code is OK.
func (r ResponseCheckTx) IsOK() bool {
	return r.Response.IsOK()
}

// IsErr returns true if Code is something other than OK.
func (r ResponseCheckTx) IsErr() bool {
	return r.Response.IsErr()
}

func (r ResponseInitChain) IsOK() bool {
	return r.Response.IsOK()
}

// IsErr returns true if Code is something other than OK.
func (r ResponseInitChain) IsErr() bool {
	return r.Response.IsErr()
}

// IsErr returns true if Code is something other than OK.
func (r ResponseBeginBlock) IsErr() bool {
	return r.Response.IsErr()
}

// IsOK returns true if Code is OK.
func (r ResponseQuery) IsOK() bool {
	return r.Code == CodeTypeOK
}

// IsErr returns true if Code is something other than OK.
func (r ResponseQuery) IsErr() bool {
	return r.Code != CodeTypeOK
}

// IsErr returns true if Code is something other than OK.
func (r ResponseEndBlock) IsErr() bool {
	return r.Response.IsErr()
}

// IsOK returns true if Code is OK.
func (r ResponseGetTxValidInfo) IsOK() bool {
	return r.Response.IsOK()
}

// IsErr returns true if Code is something other than OK.
func (r ResponseGetTxValidInfo) IsErr() bool {
	return r.Response.IsErr()
}

func (r ResponseStatus) GetMsg() string {
	return fmt.Sprintf("response code: %v, log: %v", r.Code, r.Log)
}
