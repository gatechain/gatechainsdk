package auth

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gatechain/gatechainsdk/gatechain/types"
)

func ConvertTxHashPrefixForTxResponse(txResponse types.TxResponse) types.TxResponse {
	txType := ""

	if txResponse.Logs != nil && len(txResponse.Logs) != 0 {
		if len(txResponse.Logs) > 1 {
			txType = "mix"
		} else {
			// response tag have tag-action
			for i := range txResponse.Logs {
				if txResponse.Logs[i].Events != nil {
					for k := range txResponse.Logs[i].Events {
						if txResponse.Logs[i].Events[k].Type == "message" {
							for j := range txResponse.Logs[i].Events[k].Attributes {
								if txResponse.Logs[i].Events[k].Attributes[j].Key == "action" {
									txType = txResponse.Logs[i].Events[k].Attributes[j].Value
								}
							}
							break
						}
					}
				}
			}
		}
	} else if txResponse.Tx != nil {
		if len(txResponse.Tx.GetMsgs()) == 1 {
			txType = txResponse.Tx.GetMsgs()[0].Type()
		} else if len(txResponse.Tx.GetMsgs()) > 1 {
			txType = "mix"
		}
	}

	txResponse.TxHash = AddPrefixToString(txResponse.TxHash, ChoosePrefixByType(txType))
	return txResponse
}

func GetRevocableTxSender(txResponse types.TxResponse) string {
	sender := ""
	isRevocable := false
	if txResponse.Logs != nil && len(txResponse.Logs) != 0 {
		// response tag have tag-action
		for i := range txResponse.Logs {
			if txResponse.Logs[i].Events != nil {
				for k := range txResponse.Logs[i].Events {
					if txResponse.Logs[i].Events[k].Type == "message" {
						for j := range txResponse.Logs[i].Events[k].Attributes {
							if txResponse.Logs[i].Events[k].Attributes[j].Key == "action" {
								if "revocable" == txResponse.Logs[i].Events[k].Attributes[j].Value {
									isRevocable = true
								}
							} else if txResponse.Logs[i].Events[k].Attributes[j].Key == "sender" {
								sender = txResponse.Logs[i].Events[k].Attributes[j].Value
							}
						}
						break
					}
				}
			}
		}
	}

	if isRevocable {
		if sender == "" {
			//all of this for solve the dirty tx data
			sender = txResponse.Tx.GetMsgs()[0].GetSigners()[0].String()
			events := txResponse.Events
			for i := range events {
				if events[i].Type == types.EventTypeMessage {
					attr := make([]types.Attribute, 0)
					attr = append(attr, types.NewAttribute("sender", sender))
					for _, attribute := range events[i].Attributes {
						attr = append(attr, attribute)
					}
					events[i].Attributes = attr
				}
			}
			logs := txResponse.Logs
			for i := range logs {
				if logs[i].Events != nil {
					logEvents := logs[i].Events
					for i := range logEvents {
						if logEvents[i].Type == types.EventTypeMessage {
							attr := make([]types.Attribute, 0)
							attr = append(attr, types.NewAttribute("sender", sender))
							for _, attribute := range logEvents[i].Attributes {
								attr = append(attr, attribute)
							}
							logEvents[i].Attributes = attr
						}
					}
				}
			}
		}
		return sender
	}

	return ""
}

func ChoosePrefixByType(txType string) (prefix string) {
	switch txType {
	case "send":
		prefix = "IRREVOCABLEPAY" + "-"
	case "vault":
		prefix = "VAULTCREATE" + "-"
	case "revocable":
		prefix = "REVOCABLEPAY" + "-"
	case "revoke":
		prefix = "REVOKE" + "-"
	case "updateclearingheight":
		prefix = "ACCOUNTSET" + "-"
	case "clear":
		prefix = "VAULTCLEAR" + "-"
	case "mix":
		prefix = "MIX" + "-"
	case "ethereum":
		prefix = "CONTRACT" + "-"
	default:
		prefix = "BASIC" + "-"
	}
	return prefix
}

func SplitTxHashPreFix(data string) (string, error) {
	strs := strings.Split(data, "-")
	if len(strs) == 1 {
		// no prefix
		return strs[0], nil
	} else if len(strs) != 2 {
		return data, fmt.Errorf("tx hash(%s) format error", data)
	} else {
		switch strs[0] {
		case "IRREVOCABLEPAY":
		case "VAULTCREATE":
		case "REVOCABLEPAY":
		case "REVOKE":
		case "ACCOUNTSET":
		case "VAULTCLEAR":
		case "BASIC":
		case "MIX":
		default:
			return data, fmt.Errorf("unknown tx hash(%s) prefix", data)
		}
		return strs[1], nil
	}
}

func AddPrefixToString(data, prefix string) string {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(data)
	return buffer.String()
}
