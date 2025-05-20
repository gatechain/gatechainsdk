package context

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"

	errors_pkg "github.com/pkg/errors"

	"github.com/gatechain/gatechainsdk/common"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/crypto"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/client"
	"github.com/gatechain/gatechainsdk/gatechain/types"
)

type NodeVaultQuerierImpl struct {
	Codec          *codec.Codec
	Client         *client.RestClient
	Output         io.Writer
	OutputFileName string
	OutputFormat   string
	Indent         bool
	FromAddress    types.AccAddress
	FromName       string
	Height         int64
	GenerateOnly   bool
	TrustNode      bool
	RootDir        string
}

func NewNodeVaultQuerierImpl(endpoint, API_TOKEN string) *NodeVaultQuerierImpl {
	clientUrl, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	restClient := client.MakeRestClient(*clientUrl, API_TOKEN)

	return &NodeVaultQuerierImpl{
		Client:       &restClient,
		Output:       os.Stdout,
		OutputFormat: crypto.OutputFormatText,
		Indent:       false,
	}
}

func NewNodeVaultQuerierImplGenOnly(genOnly bool, fileName, endpoint, API_TOKEN string) *NodeVaultQuerierImpl {
	clientUrl, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	restClient := client.MakeRestClient(*clientUrl, API_TOKEN)

	return &NodeVaultQuerierImpl{
		Client: &restClient,

		OutputFormat:   crypto.OutputFormatText,
		Indent:         false,
		GenerateOnly:   genOnly,
		Output:         os.Stdout,
		OutputFileName: fileName,
	}
}

func NewCLIContextWithFrom(from, endpoint, API_TOKEN, rootDir string) *NodeVaultQuerierImpl {
	clientUrl, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	restClient := client.MakeRestClient(*clientUrl, API_TOKEN)
	fromAddress, fromName, err := GetFromFields(from, false)
	if err != nil {
		fmt.Printf("failed to get from fields: %v", err)
		os.Exit(1)
	}
	return &NodeVaultQuerierImpl{
		Client:       &restClient,
		Output:       os.Stdout,
		OutputFormat: crypto.OutputFormatText,
		Indent:       false,
		FromAddress:  fromAddress,
		FromName:     fromName,
		RootDir:      rootDir,
	}
}

func (ctx *NodeVaultQuerierImpl) QueryWithData(path string, data []byte) ([]byte, int64, error) {
	return ctx.query(path, data)
}

// WithCodec returns a copy of the context with an updated codec.
func (ctx *NodeVaultQuerierImpl) WithCodec(cdc *codec.Codec) *NodeVaultQuerierImpl {
	ctx.Codec = cdc
	return ctx
}

func (ctx *NodeVaultQuerierImpl) GetCodec() *codec.Codec {
	return ctx.Codec
}

func (ctx *NodeVaultQuerierImpl) GetNode() (*client.RestClient, error) {
	if ctx.Client == nil {
		return nil, errors.New("no http client defined")
	}
	return ctx.Client, nil
}

// get the current blockchain height
func (ctx *NodeVaultQuerierImpl) GetChainHeight() (int64, error) {
	node, err := ctx.GetNode()
	if err != nil {
		return -1, err
	}
	status, err := node.Status()
	if err != nil {
		return -1, err
	}

	height := int64(status.LastRound)
	return height, nil
}

func (ctx *NodeVaultQuerierImpl) query(path string, key []byte) (res []byte, height int64, err error) {
	node, err := ctx.GetNode()
	if err != nil {
		return res, height, err
	}

	opts := appinterface.QueryOptions{
		Height: ctx.Height,
		Prove:  !ctx.TrustNode,
	}

	result, err := node.Query(path, key, opts)
	if err != nil {
		return res, height, err
	}

	resp := result.Response
	if !types.CodeType(resp.Code).IsOK() {
		return res, resp.Height, errors.New(resp.Log)
	}

	return resp.Value, resp.Height, nil
}

func (ctx *NodeVaultQuerierImpl) PrintOutput(toPrint fmt.Stringer) (err error) {
	var out []byte

	switch ctx.OutputFormat {
	case crypto.OutputFormatText:
		//out, err = yaml.Marshal(&toPrint)
		out = []byte(toPrint.String())

	case crypto.OutputFormatJSON:
		if ctx.Indent {
			out, err = ctx.Codec.MarshalJSONIndent(toPrint, "", "  ")
			out, err = AppendPrefixAddress(out, *ctx)
		} else {
			out, err = ctx.Codec.MarshalJSON(toPrint)
			out, err = AppendPrefixAddress(out, *ctx)
		}
	}

	if err != nil {
		return
	}

	fmt.Println(string(out))
	return
}

// WithHeight returns a copy of the context with an updated height.
func (ctx *NodeVaultQuerierImpl) WithHeight(height int64) *NodeVaultQuerierImpl {
	ctx.Height = height
	return ctx
}

// GetFromAddress returns the from address from the context's name.
func (ctx *NodeVaultQuerierImpl) GetFromAddress() types.AccAddress {
	return ctx.FromAddress
}

// WithFromAddress returns a copy of the context with an updated from account
// address.
func (ctx *NodeVaultQuerierImpl) WithFromAddress(addr types.AccAddress) *NodeVaultQuerierImpl {
	ctx.FromAddress = addr
	return ctx
}

// GetFromFields returns a from account address and Keybase name given either
// an address or key name. If genOnly is true, only a valid Bech32
// address is returned.
func GetFromFields(from string, genOnly bool) (types.AccAddress, string, error) {
	if from == "" {
		return nil, "", nil
	}

	if genOnly {
		addr, _, err := types.AccAddressTypeFromBech32(from)
		if err != nil {
			return nil, "", errors_pkg.Wrap(err, "must provide a valid Bech32 address for generate-only")
		}

		return addr, "", nil
	}

	keybase, err := crypto.NewKeyBaseFromDir(common.RootDir)
	if err != nil {
		return nil, "", err
	}

	var info keys.Info
	if addr, _, err := types.AccAddressTypeFromBech32(from); err == nil {
		info, err = keybase.GetByAddress(addr)
		if err != nil {
			return nil, "", err
		}
	} else {
		info, err = keybase.Get(from)
		if err != nil {
			return nil, "", err
		}
	}

	return info.GetAddress(), info.GetName(), nil
}

// GetFromName returns the key name for the current context.
func (ctx *NodeVaultQuerierImpl) GetFromName() string {
	return ctx.FromName
}
