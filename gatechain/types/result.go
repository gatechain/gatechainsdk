package types

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/spec/v1"
)

// Result is the union of ResponseFormat and ResponseCheckTx.
type Result struct {
	// Code is the response code, is stored back on the chain.
	Code CodeType

	// Codespace is the string referring to the domain of an error
	Codespace CodespaceType

	// Data is any data returned from the app.
	// Data has to be length prefixed in order to separate
	// results from multiple msgs executions
	Data []byte

	// Log contains the txs log information. NOTE: nondeterministic.
	Log string

	// GasWanted is the maximum units of work we allow this tx to perform.
	GasWanted uint64

	// GasUsed is the amount of gas actually consumed. NOTE: unimplemented
	GasUsed uint64

	// Events contains a slice of Event objects that were emitted during some
	// execution.
	Events Events
}

const (
	// the irrevocablepay tx
	IRREVOCABLEPAY = uint8(0)
	// the revocablepay tx
	REVOCABLEPAY = uint8(1)
	// the revoked tx
	REVOKED = uint8(2)
)

func RevocableStatusStr(status uint8) (str string) {
	switch status {
	case IRREVOCABLEPAY:
		str = "IRREVOCABLEPAY"
	case REVOCABLEPAY:
		str = "REVOCABLEPAY"
	case REVOKED:
		str = "REVOKED"
	}
	return str
}

// ABCIMessageLogs represents a slice of ABCIMessageLog.
type ABCIMessageLogs []ABCIMessageLog

// ABCIMessageLog defines a structure containing an indexed tx ABCI message log.
type ABCIMessageLog struct {
	MsgIndex uint16 `json:"msg_index"`
	Success  bool   `json:"success"`
	Log      string `json:"log"`

	// Events contains a slice of Event objects that were emitted during some
	// execution.
	Events StringEvents `json:"events"`
}

// String implements the fmt.Stringer interface for the ABCIMessageLogs type.
func (logs ABCIMessageLogs) String() (str string) {
	if logs != nil {
		raw, err := codec.Cdc.MarshalJSON(logs)
		if err == nil {
			str = string(raw)
		}
	}

	return str
}

// TxResponse defines a structure containing relevant tx data and metadata. The
// tags are stringified and the log is JSON decoded.
type TxResponse struct {
	Height                   int64           `json:"height"`
	TxHash                   string          `json:"txhash"`
	Code                     uint32          `json:"code,omitempty"`
	Data                     string          `json:"data,omitempty"`
	RawLog                   string          `json:"raw_log,omitempty"`
	Logs                     ABCIMessageLogs `json:"logs,omitempty"`
	Info                     string          `json:"info,omitempty"`
	GasWanted                int64           `json:"gas_wanted,omitempty"`
	GasUsed                  int64           `json:"gas_used,omitempty"`
	Codespace                string          `json:"codespace,omitempty"`
	Tx                       Tx              `json:"tx,omitempty"`
	Timestamp                string          `json:"timestamp,omitempty"`
	IrrevocableConfirmNumber []int64         `json:"irrevocable_confirm_number"`

	// DEPRECATED: Remove in the next next major release in favor of using the
	// ABCIMessageLog.Events field.
	Events StringEvents `json:"events,omitempty"`
}

// RevocableTxResponse defines a structure containing revocable tx data and metadata
type RevocableTxResponse struct {
	Status       uint8  `json:"status"`
	RevokeTxHash string `json:"revoke_hash"`
}

// NewResponseResultTx returns a TxResponse given a ResultTx from tendermint
func NewRevocableTxResponse(status uint8, revokeTxHash string) *RevocableTxResponse {
	return &RevocableTxResponse{
		Status:       status,
		RevokeTxHash: revokeTxHash,
	}
}

func (r RevocableTxResponse) String() string {
	var sb strings.Builder
	sb.WriteString("Response:\n")

	sb.WriteString(fmt.Sprintf("  Status: %s\n", RevocableStatusStr(r.Status)))

	if r.RevokeTxHash != "" {
		sb.WriteString(fmt.Sprintf("  RevokeTxHash: %s\n", r.RevokeTxHash))
	}

	return strings.TrimSpace(sb.String())
}

// NewResponseResultTx returns a TxResponse given a ResultTx from tendermint
func NewResponseResultTx(res *appinterface.ResponseTx, tx Tx, timestamp string) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	parsedLogs, _ := ParseABCILogs(res.Response.Log)
	return TxResponse{
		TxHash: fmt.Sprintf("%X", appinterface.Tx(res.Tx).Hash()),
		Height: res.Height,
		Code:   res.Response.Code,
		Data:   strings.ToUpper(hex.EncodeToString(res.Tx)),
		RawLog: res.Response.Log,
		Logs:   parsedLogs,
		//Info:      res.Response.Info,
		GasWanted: int64(res.Response.GasWanted),
		GasUsed:   int64(res.Response.GasUsed),
		Events:    StringifyEvents(res.Response.Events),
		Tx:        tx,
		Timestamp: timestamp,
	}
}

// NewResponseFormatBroadcastTx returns a TxResponse given a ResultBroadcastTx from tendermint
func NewResponseFormatBroadcastTx(res *v1.ResultBroadcastTx) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	parsedLogs, _ := ParseABCILogs(res.Log)

	var txHashStr string

	txHashStr = fmt.Sprintf("%X", res.Hash)

	return TxResponse{
		Code:   res.Code,
		Data:   base64.StdEncoding.EncodeToString(res.Data),
		RawLog: res.Log,
		Logs:   parsedLogs,
		TxHash: txHashStr,
	}
}

func (r TxResponse) String() string {
	var sb strings.Builder
	sb.WriteString("Response:\n")

	if r.Height > 0 {
		sb.WriteString(fmt.Sprintf("  Height: %d\n", r.Height))
	}

	if r.TxHash != "" {
		sb.WriteString(fmt.Sprintf("  TxHash: %s\n", r.TxHash))
	}

	if r.Code > 0 {
		sb.WriteString(fmt.Sprintf("  Code: %d\n", r.Code))
	}

	if r.Data != "" {
		sb.WriteString(fmt.Sprintf("  Data: %s\n", r.Data))
	}

	if r.RawLog != "" {
		sb.WriteString(fmt.Sprintf("  Raw Log: %s\n", r.RawLog))
	}

	if r.Logs != nil {
		sb.WriteString(fmt.Sprintf("  Logs: %s\n", r.Logs))
	}

	if r.Info != "" {
		sb.WriteString(fmt.Sprintf("  Info: %s\n", r.Info))
	}

	if r.GasWanted != 0 {
		sb.WriteString(fmt.Sprintf("  GasWanted: %d\n", r.GasWanted))
	}

	if r.GasUsed != 0 {
		sb.WriteString(fmt.Sprintf("  GasUsed: %d\n", r.GasUsed))
	}

	if r.Codespace != "" {
		sb.WriteString(fmt.Sprintf("  Codespace: %s\n", r.Codespace))
	}

	if r.Timestamp != "" {
		sb.WriteString(fmt.Sprintf("  Timestamp: %s\n", r.Timestamp))
	}

	if len(r.Events) > 0 {
		sb.WriteString(fmt.Sprintf("  Events: \n%s\n", r.Events.String()))
	}

	return strings.TrimSpace(sb.String())
}

// Empty returns true if the response is empty
func (r TxResponse) Empty() bool {
	return r.TxHash == "" && r.Logs == nil
}

// ParseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// ABCIMessageLog types. It returns an error upon JSON decoding failure.
func ParseABCILogs(logs string) (res ABCIMessageLogs, err error) {
	err = json.Unmarshal([]byte(logs), &res)
	return res, err
}
