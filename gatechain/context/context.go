package context

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"

	errors2 "github.com/pkg/errors"

	"github.com/gatechain/gatechainsdk/api/common"
	"github.com/gatechain/gatechainsdk/gatechain/codec"
	"github.com/gatechain/gatechainsdk/gatechain/crypto"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
	"github.com/gatechain/gatechainsdk/gatechain/node/appinterface"
	"github.com/gatechain/gatechainsdk/gatechain/rpc/client"
	types2 "github.com/gatechain/gatechainsdk/gatechain/types"
)

type NodeVaultQuerierImpl struct {
	Codec          *codec.Codec
	Client         *client.RestClient
	Output         io.Writer
	OutputFileName string
	OutputFormat   string
	Indent         bool
	FromAddress    types2.AccAddress
	FromName       string
	Height         int64
	GenerateOnly   bool
	//NodeURI               string
	//From                  string
	TrustNode bool
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
		OutputFormat: "text", //"text"  "json"
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

		OutputFormat:   "text", //"text"  "json"
		Indent:         false,
		GenerateOnly:   genOnly,
		Output:         os.Stdout,
		OutputFileName: fileName,
	}
}
func NewCLIContextWithFrom(from, endpoint, API_TOKEN string) *NodeVaultQuerierImpl {
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
		OutputFormat: "text", //"text"  "json"
		Indent:       false,
		FromAddress:  fromAddress,
		FromName:     fromName,
	}
}

func (ctx *NodeVaultQuerierImpl) QueryWithData(path string, data []byte) ([]byte, int64, error) {
	return ctx.query(path, data)
}

// QueryStore performs a query to a Tendermint node with the provided key and
// store name. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx *NodeVaultQuerierImpl) QueryStore(key HexBytes, storeName string) ([]byte, int64, error) {
	return ctx.queryStore(key, storeName, "key")
}

// queryStore performs a query to a Tendermint node with the provided a store
// name and path. It returns the result and height of the query upon success
// or an error if the query fails.
func (ctx *NodeVaultQuerierImpl) queryStore(key HexBytes, storeName, endPath string) ([]byte, int64, error) {
	path := fmt.Sprintf("/store/%s/%s", storeName, endPath)
	return ctx.query(path, key)
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
	if !types2.CodeType(resp.Code).IsOK() {
		return res, resp.Height, errors.New(resp.Log)
	}

	return resp.Value, resp.Height, nil
}

func (ctx *NodeVaultQuerierImpl) PrintOutput(toPrint fmt.Stringer) (err error) {
	var out []byte

	switch ctx.OutputFormat {
	case "text":
		//out, err = yaml.Marshal(&toPrint)
		out = []byte(toPrint.String())

	case "json":
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
func (ctx *NodeVaultQuerierImpl) GetFromAddress() types2.AccAddress {
	return ctx.FromAddress
}

// WithFromAddress returns a copy of the context with an updated from account
// address.
func (ctx *NodeVaultQuerierImpl) WithFromAddress(addr types2.AccAddress) *NodeVaultQuerierImpl {
	ctx.FromAddress = addr
	return ctx
}

// GetFromFields returns a from account address and Keybase name given either
// an address or key name. If genOnly is true, only a valid Bech32
// address is returned.
func GetFromFields(from string, genOnly bool) (types2.AccAddress, string, error) {
	if from == "" {
		return nil, "", nil
	}

	if genOnly {
		addr, _, err := types2.AccAddressTypeFromBech32(from)
		if err != nil {
			return nil, "", errors2.Wrap(err, "must provide a valid Bech32 address for generate-only")
		}

		return addr, "", nil
	}

	keybase, err := crypto.NewKeyBaseFromDir(common.RootDir)
	if err != nil {
		return nil, "", err
	}

	var info keys.Info
	if addr, _, err := types2.AccAddressTypeFromBech32(from); err == nil {
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
