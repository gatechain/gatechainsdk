module github.com/gatechain/gatechainsdk

go 1.23.4

require (
	cosmossdk.io/errors v1.0.2
	github.com/cosmos/go-bip39 v1.0.0
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c
	github.com/gatechain/crypto v0.1.0
	github.com/google/go-querystring v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.4.0
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tm-db v0.6.7
	golang.org/x/crypto v0.32.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/ChainSafe/go-schnorrkel v1.0.0 // indirect
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cosmos/gorocksdb v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/ethereum/go-ethereum v1.10.8 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gatechain/go-deadlock v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.2.4 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/klauspost/compress v1.17.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/maoxs2/go-ripemd v0.0.0-20200227063901-215d858ac634 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20210601165009-122bf33a46e0 // indirect
	github.com/onsi/gomega v1.20.0 // indirect
	github.com/petermattis/goid v0.0.0-20231207134359-e60b3f734c67 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
	google.golang.org/grpc v1.71.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace (
	github.com/btcsuite/btcd/btcec => github.com/btcsuite/btcd/btcec/v2 v2.3.4
	github.com/cosmos/cosmos-sdk => github.com/celestiaorg/cosmos-sdk v1.23.0-sdk-v0.46.16
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/celestiaorg/celestia-core v1.38.0-tm-v0.34.29
)

exclude github.com/btcsuite/btcd/btcec v0.0.0-00010101000000-000000000000
