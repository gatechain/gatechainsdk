package context

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/gatechain/gatechainsdk/gatechain/types"
)

// gatechain add
// AppendPrefixAddress returns byte that accAddress append different prefix.
func AppendPrefixAddress(buffer []byte, cliCtx NodeVaultQuerierImpl) ([]byte, error) {
	var data interface{}

	err := json.Unmarshal(buffer, &data)
	if err != nil {
		return nil, err
	}
	//unmarshal data to map[string]interface{}
	switch vv := data.(type) {
	case []interface{}:
		newValue := make([]interface{}, 0)
		for i := range vv {
			switch vv[i].(type) {
			case string:
				if isAccAddress(vv[i].(string)) {
					addr, _, err := sdk.AccAddressTypeFromBech32(vv[i].(string))
					if err != nil {
						return nil, err
					}
					vv[i], err = ConvertAddress(cliCtx, addr)
					if err != nil {
						return nil, err
					}
				}
			}
			subBuffer, err := json.Marshal(vv[i])
			if err != nil {
				return nil, err
			}
			newSubBuffer, err := AppendPrefixAddress(subBuffer, cliCtx)
			if err != nil {
				return nil, err
			}
			var (
				subData interface{}
			)
			if err := json.Unmarshal(newSubBuffer, &subData); err != nil {
				return nil, err
			}
			newValue = append(newValue, subData)
		}
		data = newValue
	case map[string]interface{}:
		newValue := make(map[string]interface{})
		for k := range vv {
			switch vv[k].(type) {
			case string:
				// key whether contains "address" and value is type of accaddress, modify address-prefix
				if isAccAddress(vv[k].(string)) {
					addr, _, err := sdk.AccAddressTypeFromBech32(vv[k].(string))
					if err != nil {
						return nil, err
					}
					vv[k], err = ConvertAddress(cliCtx, addr)
					if err != nil {
						return nil, err
					}
				}
			}
			// convert map[string]interface{}
			subBuffer, err := json.Marshal(vv[k])
			if err != nil {
				return nil, err
			}
			newSubBuffer, err := AppendPrefixAddress(subBuffer, cliCtx)
			if err != nil {
				return nil, err
			}
			var subData interface{}
			if err := json.Unmarshal(newSubBuffer, &subData); err != nil {
				return nil, err
			}
			newValue[k] = subData
		}
		data = newValue
	}
	bz, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

// isAccAddress check address is type of AccAddress.
func isAccAddress(address string) bool {
	_, _, err := sdk.AccAddressTypeFromBech32(address)
	return err == nil
}

func ConvertAddress(cliCtx NodeVaultQuerierImpl, addr sdk.AccAddress) (string, error) {
	accountType, err := GetVaultAccountType(cliCtx, addr)
	if err != nil {
		return "", err
	}
	return addr.TypeString(accountType), err
}

func GetVaultAccountType(cliCtx NodeVaultQuerierImpl, addr sdk.AccAddress) (uint8, error) {
	bs, err := cliCtx.Codec.MarshalJSON(addr)
	if err != nil {
		return uint8(0), nil
	}
	// query vault account and allow not to found
	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/acc/check"), bs)
	accountType, err := strconv.Atoi(string(res))
	if err != nil {
		return uint8(0), nil
	}
	return uint8(accountType), nil
}
