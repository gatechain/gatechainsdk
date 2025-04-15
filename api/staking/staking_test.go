package staking

import (
	"fmt"
	"testing"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

// done
// Query the delegation information of a delegated account under a single consensus account.
// gatecli staking delegation [delegator-addr] [con-account-addr]
// gatecli staking delegation gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04 gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
func TestGetCmdQueryDelegation(t *testing.T) {
	delegatorAddr := "gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04"
	validatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	GetCmdQueryDelegation(ctx, common.Staking, delegatorAddr, validatorAddr)
}

// done
/*
Query the delegation information of a delegated account across all consensus accounts.
Use:   "delegations [delegator-addr]",
Short: "Query all delegations made by one delegator",

gatecli staking delegations gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
Delegation:
  Delegator:   gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
  Con-account: gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
  Shares:      300000000000000000000000000.000000000000000000
  Balance:   300000000000000000000000000

gatecli staking delegations gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04

gatecli staking delegations gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
Delegation:
  Delegator:   gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
  Con-account: gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
  Shares:      100000000.000000000000000000
  Balance:   100000000
Delegation:
  Delegator:   gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
  Con-account: gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
  Shares:      626000280.000000000000000000
  Balance:   626000280


gatecli staking delegations gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
*/
func TestGetCmdQueryDelegations(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	delegatorAddr := "gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04"
	GetCmdQueryDelegations(ctx, common.Staking, delegatorAddr)
}

// done
//
/*
Query all undelegated (unbonded) delegations from the specified consensus account.

gatecli staking undelegations-from gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz
--node http://139.162.15.16:8080 --home /mnt/meteora-data/
*/
func TestGetCmdQueryValidatorUnbondingDelegations(t *testing.T) {
	ValidatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	GetCmdQueryValidatorUnbondingDelegations(ctx, common.Staking, ValidatorAddr)
}

// done
// gatecli staking pool
/*
gatecli staking pool
Pool:
  Not Bonded Tokens:  0
  Bonded Tokens:      300000000000000001452000560
*/
func TestGetCmdQueryPool(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	GetCmdQueryPool(ctx, common.Staking)
}

// done
// gatecli staking params
/*
 gatecli staking params
Params:
  Undelegating Time:    10m0s
  Max Con-accounts:    100
  Max Entries:       7
  Bonded Coin Denom: NANOGT
  Pow Rate:          1
  Max Pow Rate:      2
*/
func TestGetCmdQueryParams(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	GetCmdQueryParams(ctx, common.Staking)
}

// done
// Query all redelegation records of the delegator account.
// Use:   "redelegations [delegator-addr]",
// gatecli staking redelegations  gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
/*
Redelegations between:
  Delegator:                   gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
  Source Con-account:          gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
  Destination Con-account:     gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
  Entries:
    Redelegation Entry #0:
      Creation height:           8944
      Min time to undelegate (unix): 2025-05-14 10:45:06 +0000 UTC
      Initial Balance:           100
      Shares:                    100.000000000000000000
      Balance:                   100
*/
func TestGetCmdQueryRedelegations(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	delegatorAddr := "gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04"
	GetCmdQueryRedelegations(ctx, common.Staking, delegatorAddr)
}

// todo
/*
Use:   "redelegation [delegator-addr] [src-validator-addr] [dst-validator-addr]",
Short: "Query a redelegation record based on delegator and a source and destination validator address",

Query redelegation records of the delegator account between two validator accounts.

Redelegations between:
  Delegator:                   gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
  Source Con-account:          gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
  Destination Con-account:     gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
  Entries:
    Redelegation Entry #0:
      Creation height:           8944
      Min time to undelegate (unix): 2025-05-14 10:45:06 +0000 UTC
      Initial Balance:           100
      Shares:                    100.000000000000000000
      Balance:                   100
*/
// gatecli staking redelegation gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
// gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
// gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9

func TestGetCmdQueryRedelegation(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	args := []string{
		"gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04",
		"gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh",
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9",
	}
	GetCmdQueryRedelegation(ctx, common.Staking, args)
}

// done
// gatecli staking delegations-to gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
func TestGetCmdQueryValidatorDelegations(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	delegatorAddr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	GetCmdQueryValidatorDelegations(ctx, common.Staking, delegatorAddr)
}

//done
// gatecli staking undelegation [delegator-addr] [con-account_addr]
// gatecli staking undelegation gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
/*
gatecli staking undelegation gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
Undelegations  between:
  Delegator:   gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
  Con-account: gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
	Entries:
		Undelegation         0:
		Creation Height:           8908
		Min time to undelegate (unix): 2025-05-14 10:27:30 +0000 UTC
		Expected balance:          100
*/
func TestGetCmdQueryUnbondingDelegation(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	delegatorAddr := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	validatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	GetCmdQueryUnbondingDelegation(ctx, common.Staking, delegatorAddr, validatorAddr)
}

/* done
Use:   "undelegations [delegator-addr]",
		Short: "Query all undelegations records for one delegator",
gatecli staking undelegations gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz


gatecli staking undelegations gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
Undelegations  between:
  Delegator:   gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
  Con-account: gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
	Entries:
		Undelegation         0:
		Creation Height:           8908
		Min time to undelegate (unix): 2025-05-14 10:27:30 +0000 UTC
		Expected balance:          100
*/

func TestGetCmdQueryUnbondingDelegations(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	delegatorAddr := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	GetCmdQueryUnbondingDelegations(ctx, common.Staking, delegatorAddr)
}

//gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
//done
/*
	Use:   "redelegations-from [con-account]",
    Short: "Query all outgoing redelegatations from a con-account",

Redelegations between:
  Delegator:                   gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz
  Source Con-account:          gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
  Destination Con-account:     gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
  Entries:
    Redelegation Entry #0:
      Creation height:           8856
      Min time to undelegate (unix): 2025-05-14 10:01:40 +0000 UTC
      Initial Balance:           100
      Shares:                    100.000000000000000000
      Balance:                   100
*/
func TestGetCmdQueryValidatorRedelegations(t *testing.T) {
	cdc := utils.MakeCodec()
	ctx := context.NewNodeVaultQuerierImpl(common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)
	SrcValidatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"

	GetCmdQueryValidatorRedelegations(ctx, common.Staking, SrcValidatorAddr)
}

/*
gatecli staking delegate gt11pqt5ugs9vwftjmwpyvgzn98jtquxzqrmrgqehl3p7w3k4np88m2lk4mr0nhqprqz93egxa
20000000000NANOGT --chain-id gate-66
--from gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4  --fees 1000000NANOGT --yes
*/

// delegate gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4  to
// con-account gt11pqt5ugs9vwftjmwpyvgzn98jtquxzqrmrgqehl3p7w3k4np88m2lk4mr0nhqprqz93egxa
/*

 gatecli staking delegate gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh
363000140NANOGT --chain-id gate-66 --from gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04
--fees 100000000NANOGT -y

*/
func TestGetCmdDelegate(t *testing.T) {
	cdc := utils.MakeCodec()
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	ctx := context.NewCLIContextWithFrom(delegatorAddress, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	txBldr := auth2.NewTxBuilderFromCLI()
	txBldr = txBldr.WithTxEncoder(types.DefaultTxEncoder(cdc))
	txBldr = txBldr.WithFees("100000000NANOGT")
	txBldr = txBldr.WithGas(200000)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	txBldr = txBldr.WithChainID("gate-66")
	fmt.Println(txBldr, err)

	amountStr := "363000140NANOGT"
	validatorAddr := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh" //args[0]

	GetCmdDelegate(ctx, txBldr, amountStr, delegatorAddress, validatorAddr)
}

// gatecli staking  redelegate [src-con-account] [dst-con-account] [amount]
func TestGetCmdRedelegate(t *testing.T) {
	cdc := utils.MakeCodec()
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	ctx := context.NewCLIContextWithFrom(delegatorAddress, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	txBldr := auth2.NewTxBuilderFromCLI()
	txBldr = txBldr.WithTxEncoder(types.DefaultTxEncoder(cdc))
	txBldr = txBldr.WithFees("100000000NANOGT")
	txBldr = txBldr.WithGas(410000)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	txBldr = txBldr.WithChainID("gate-66")
	fmt.Println(txBldr, err)

	args := []string{
		"gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh",
		"gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9",
		"100NANOGT",
	}
	GetCmdRedelegate(ctx, txBldr, delegatorAddress, args)
}

//done 2
// unbond from consensus account
// gt11hqsern6mu8mfykjs9drchh227pfr36w9m6md3dhxtg382kwk39njm69qczmzj49sxqxl04  gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9
//
/*
gatecli staking undelegate gt11s2pn5kwec3kc3aw86u40vf73dt3gar24qqzjy00ezmeracw497u79g6ykjxgpzjmmwelan 549262360NANOGT
--chain-id meteora --from gt11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq99na4f
--fees 10000000NANOGT --node http://139.162.15.16:8080 --home /mnt/meteora-data/ -y
*/
func TestGetCmdUnbond(t *testing.T) {
	cdc := utils.MakeCodec()
	delegatorAddress := "gt112ls6cyr86e6zyd9akcagcegeqygjqwt3cqp9w3kk48jtepel2k8tjt5ggwgp36887735lz"
	ctx := context.NewCLIContextWithFrom(delegatorAddress, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	txBldr := auth2.NewTxBuilderFromCLI()
	txBldr = txBldr.WithTxEncoder(types.DefaultTxEncoder(cdc))
	txBldr = txBldr.WithFees("100000000NANOGT")
	txBldr = txBldr.WithGas(200000)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	txBldr = txBldr.WithChainID("gate-66")
	fmt.Println(txBldr, err)

	validatorAddress := "gt1147zmtrfu3w4qn3d7qtlrm0t9j8h9t2pg4e8cldndn22ajca8g3jhrq0sl2kc4m5tthravh"
	amountStr := "100NANOGT"

	GetCmdUnbond(ctx, txBldr, amountStr, delegatorAddress, validatorAddress)
}

// Proxy undelegation
/*
gatecli staking undelegate-by-retrieval-account vault11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq49935w
--from gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz --gas 250000 --fees 100000000NANOGT
--chain-id gate-66
*/
func TestGetCmdUnbondBySecAddr(t *testing.T) {
	securityAddress := "gt11vk8xsf2zfr7yzuadpa3t08e2qju4ps4x8pp6a4rkakgw3kup0gm4jy50s5uuyf45uf6ykz"
	cdc := utils.MakeCodec()
	ctx := context.NewCLIContextWithFrom(securityAddress, common.EndPoint, common.APIToken)
	ctx = ctx.WithCodec(cdc)

	txBldr := auth2.NewTxBuilderFromCLI()
	txBldr = txBldr.WithTxEncoder(types.DefaultTxEncoder(cdc))
	txBldr = txBldr.WithFees("100000000NANOGT")
	txBldr = txBldr.WithGas(200000)
	txBldr, err := auth2.UpdateValidHeight(ctx, txBldr)
	txBldr = txBldr.WithChainID("gate-66")
	fmt.Println(txBldr, err)

	args := []string{
		"vault11mzum4y48hn8xsw3hcdjwx7mu5zxwhnds6m3fqz4wrfteg0z9rkyt9yczrrjhjalq49935w",
	}

	GetCmdUnbondBySecAddr(ctx, txBldr, args)
}
