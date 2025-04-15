package cli

import (
	"fmt"
	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/api/utils"
	auth2 "github.com/gatechain/gatechainsdk/gatechain/auth"
	"github.com/gatechain/gatechainsdk/gatechain/context"
	"github.com/gatechain/gatechainsdk/gatechain/types"
	"testing"
)

/*
gatecli staking delegate gt11pqt5ugs9vwftjmwpyvgzn98jtquxzqrmrgqehl3p7w3k4np88m2lk4mr0nhqprqz93egxa
20000000000NANOGT --chain-id gate-66
--from gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4  --fees 1000000NANOGT --yes
*/

// delegate gt11zzu8glwr7mc5t7w2tqtwuvn6c3scd7x8hs7f4ud7379r7xknkyxsx4pnhnh50v9t7ekkm4  to
// consensus account gt11pqt5ugs9vwftjmwpyvgzn98jtquxzqrmrgqehl3p7w3k4np88m2lk4mr0nhqprqz93egxa
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
//
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

//
/*
“Proxy undelegation”
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
