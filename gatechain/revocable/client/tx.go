package client

import (
	"fmt"

	"github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/auth/exported"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

func EnsureFromVaultAccount(ctx context.NodeVaultQuerierImpl) error {
	_, err := queryVaultAccount(ctx, ctx.GetFromAddress())
	return err
}

// queryVaultAccount queries an vault account using custom query endpoint of auth module
// returns an error if result is `null` otherwise account data
func queryVaultAccount(ctx context.NodeVaultQuerierImpl, addr types.AccAddress) ([]byte, error) {
	bz, err := ctx.Codec.MarshalJSON(auth.NewQueryVaultAccountParams(addr))
	if err != nil {
		return nil, err
	}

	route := fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount)

	res, _, err := ctx.QueryWithData(route, bz)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetVaultAccount(ctx context.NodeVaultQuerierImpl, address []byte) (exported.VaultAccount, error) {
	res, err := queryVaultAccount(ctx, address)
	if err != nil {
		return nil, err
	}

	var account exported.VaultAccount
	if err := ctx.Codec.UnmarshalJSON(res, &account); err != nil {
		return nil, err
	}
	return account, nil
}

func GetRevocableTokens(ctx context.NodeVaultQuerierImpl, address []byte, height int64) (exported.RevocableTxCoinsArray, error) {

	res, err := GetVaultAccount(ctx, address)
	if err != nil {
		return nil, err
	}

	coins, err := res.GetRevocableTokensDetail(height)
	if err != nil {
		return nil, err
	}
	return coins, nil
}
