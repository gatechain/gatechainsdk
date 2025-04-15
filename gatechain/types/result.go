package types

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	appinterface2 "github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/spec/v1"
	//"github.com/gatechain/gatemint/data/transactions"
	"math"
	"strings"

	//v1 "github.com/gatechain/gatemint/daemon/gmd/api/spec/v1"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
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

// SimulationResponse defines the response generated when a transaction is successfully
// simulated by the Baseapp.
type SimulationResponse struct {
	GasInfo
	Result *Result
}

// GasInfo defines tx execution gas context.
type GasInfo struct {
	// GasWanted is the maximum units of work we allow this tx to perform.
	GasWanted uint64

	// GasUsed is the amount of gas actually consumed.
	GasUsed uint64
}

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

// TODO: In the future, more codes may be OK.
func (res Result) IsOK() bool {
	return res.Code.IsOK()
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

func NewABCIMessageLog(i uint16, success bool, log string, events Events) ABCIMessageLog {
	return ABCIMessageLog{
		MsgIndex: i,
		Success:  success,
		Log:      log,
		Events:   StringifyEvents(events.ToABCIEvents()),
	}
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
func NewResponseResultTx(res *appinterface2.ResponseTx, tx Tx, timestamp string) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	parsedLogs, _ := ParseABCILogs(res.Response.Log)
	return TxResponse{
		TxHash: fmt.Sprintf("%X", appinterface2.Tx(res.Tx).Hash()),
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

// NewResponseFormatBroadcastTxCommit returns a TxResponse given a
// ResultBroadcastTxCommit from tendermint.
func NewResponseFormatBroadcastTxCommit(res *ctypes.ResultBroadcastTxCommit) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	if !res.CheckTx.IsOK() {
		return newTxResponseCheckTx(res)
	}

	return newTxResponseDeliverTx(res)
}

func newTxResponseCheckTx(res *ctypes.ResultBroadcastTxCommit) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	var txHash string
	if res.Hash != nil {
		txHash = res.Hash.String()
	}

	parsedLogs, _ := ParseABCILogs(res.CheckTx.Log)

	return TxResponse{
		Height:    res.Height,
		TxHash:    txHash,
		Code:      res.CheckTx.Code,
		Data:      strings.ToUpper(hex.EncodeToString(res.CheckTx.Data)),
		RawLog:    res.CheckTx.Log,
		Logs:      parsedLogs,
		Info:      res.CheckTx.Info,
		GasWanted: res.CheckTx.GasWanted,
		GasUsed:   res.CheckTx.GasUsed,
		//Events:    StringifyEvents(res.CheckTx.Events),
		Codespace: res.CheckTx.Codespace,
	}
}

func newTxResponseDeliverTx(res *ctypes.ResultBroadcastTxCommit) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	var txHash string
	if res.Hash != nil {
		txHash = res.Hash.String()
	}

	parsedLogs, _ := ParseABCILogs(res.DeliverTx.Log)

	return TxResponse{
		Height:    res.Height,
		TxHash:    txHash,
		Code:      res.DeliverTx.Code,
		Data:      strings.ToUpper(hex.EncodeToString(res.DeliverTx.Data)),
		RawLog:    res.DeliverTx.Log,
		Logs:      parsedLogs,
		Info:      res.DeliverTx.Info,
		GasWanted: res.DeliverTx.GasWanted,
		GasUsed:   res.DeliverTx.GasUsed,
		//Events:    StringifyEvents(res.DeliverTx.Events),
		Codespace: res.DeliverTx.Codespace,
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

// SearchTxsResult defines a structure for querying txs pageable
type SearchTxsResult struct {
	TotalCount int          `json:"total_count"` // Count of all txs
	Count      int          `json:"count"`       // Count of txs in current page
	PageNumber int          `json:"page_number"` // Index of current page, start from 1
	PageTotal  int          `json:"page_total"`  // Count of total pages
	Limit      int          `json:"limit"`       // Max count txs per page
	Txs        []TxResponse `json:"txs"`         // List of txs in current page
}

func NewSearchTxsResult(totalCount, count, page, limit int, txs []TxResponse) SearchTxsResult {
	return SearchTxsResult{
		TotalCount: totalCount,
		Count:      count,
		PageNumber: page,
		PageTotal:  int(math.Ceil(float64(totalCount) / float64(limit))),
		Limit:      limit,
		Txs:        txs,
	}
}

// ParseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// ABCIMessageLog types. It returns an error upon JSON decoding failure.
func ParseABCILogs(logs string) (res ABCIMessageLogs, err error) {
	err = json.Unmarshal([]byte(logs), &res)
	return res, err
}
