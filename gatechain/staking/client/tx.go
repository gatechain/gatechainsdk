package client

import (
	"errors"
	"fmt"
	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	types2 "github.com/gatechain/gatechainsdk/gatechain/staking/types"
	"github.com/gatechain/gatechainsdk/gatechain/types"
	"github.com/spf13/viper"
)

// BuildCreateValidatorMsg makes a new MsgCreateValidator.
func BuildCreateValidatorMsg(cliCtx context.NodeVaultQuerierImpl, txBldr auth.TxBuilder) (auth.TxBuilder, types.Msg, error) {
	valAddr := cliCtx.GetFromAddress()
	pkStr := viper.GetString(FlagPubKey)
	if pkStr == "" {
		return txBldr, nil, errors.New("must use flag pubkey")
	}

	monikerStr := viper.GetString(FlagMoniker)
	if monikerStr == "" {
		return txBldr, nil, errors.New("must use flag moniker")
	}

	pk, err := types.GetAccPubKeyBech32(pkStr)
	if err != nil {
		return txBldr, nil, err
	}

	description := types2.NewDescription(
		viper.GetString(FlagMoniker),
		viper.GetString(FlagIdentity),
		viper.GetString(FlagWebsite),
		viper.GetString(FlagDetails),
	)

	// get the initial validator commission parameters
	rateStr := viper.GetString(FlagCommissionRate)
	maxRateStr := viper.GetString(FlagCommissionMaxRate)
	maxChangeRateStr := viper.GetString(FlagCommissionMaxChangeRate)
	commissionRates, err := buildCommissionRates(rateStr, maxRateStr, maxChangeRateStr)
	if err != nil {
		return txBldr, nil, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return txBldr, nil, err
	}

	data, err := node.GetParticipationKey(types.GetAddressTO48(valAddr))
	if err != nil {
		return txBldr, nil, err
	}
	if data.Code != 00 {
		return txBldr, nil, fmt.Errorf(data.Log)
	}

	//extra := gen.MakeOnlineExtra(data.Data)//todo future
	var extra []byte
	msg := types2.NewMsgCreateValidator(
		types.ValAddress(valAddr), pk, description, commissionRates, extra)

	//if viper.GetBool(client.FlagGenerateOnly) {
	//	ip := viper.GetString(FlagIP)
	//	nodeID := viper.GetString(FlagNodeID)
	//	if nodeID != "" && ip != "" {
	//		txBldr = txBldr.WithMemo(fmt.Sprintf("%s@%s:26656", nodeID, ip))
	//	}
	//}

	return txBldr, msg, nil
}
