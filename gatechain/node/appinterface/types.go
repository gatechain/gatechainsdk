package appinterface

import (
	"fmt"
	"github.com/davidlazar/go-crypto/encoding/base32"
	"github.com/gatechain/crypto/merkle"
	"strings"

	"github.com/gatechain/gatechainsdk/gatechain/node/basics"
)

type RequestQuery struct {
	Data   []byte `json:"data"`
	Path   string `json:"path"`
	Height int64  `json:"height"`
	Prove  bool   `json:"Prove"`
}

type QueryOptions struct {
	Height int64
	Prove  bool
}
type ResponseQuery struct {
	Code uint32 `json:"code"`
	// bytes data = 2; // use "value" instead.
	Log   string `json:"log"`
	Info  string `json:"info"`
	Index int64  `json:"index"`
	Key   []byte `json:"key"`
	Value []byte `json:"value"`

	//TODO need to change to gt merkle tree
	Proof     *merkle.Proof `json:"proof"`
	Height    int64         `json:"height"`
	Codespace string        `json:"codespace"`
}

type ResponseStatus struct {
	Type      string
	Code      uint32
	Log       string
	GasWanted uint64
	GasUsed   uint64
	Events    []Event
	Data      []byte
}

type Event struct {
	Type       string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes []KVPair `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty"`
}
type KVPair struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

type ParticipationData struct {
	Address         []byte
	VoteID          []byte
	SelectionID     []byte
	VoteKeyDilution uint64
	OnlineStatus    string
}

func (pd ParticipationData) String() string {
	var sb strings.Builder
	sb.WriteString("ParticipationData:\n")
	if pd.Address != nil {
		sb.WriteString(fmt.Sprintf("  Address: 		%s\n", basics.ConverAddress(pd.Address)))
	}
	if pd.VoteID != nil {
		sb.WriteString(fmt.Sprintf("  VoteID: 		%s\n", base32.EncodeToString(pd.VoteID)))
	}
	if pd.SelectionID != nil {
		sb.WriteString(fmt.Sprintf("  SelectionID: 		%s\n", base32.EncodeToString(pd.SelectionID)))
	}
	return strings.TrimSpace(sb.String())
}

type ResponseTx struct {
	Height         int64          `json:"height"`
	Index          uint32         `json:"index"`
	Tx             []byte         `json:"tx"`
	Response       ResponseStatus `json:"response"`
	ResponseTxData []ResponseTxExtra
}
type ResponseTxExtra struct {
	Data  []byte
	Index uint64
}
