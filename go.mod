module github.com/kogisin/dessert

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/cosmos/cosmos-sdk v0.34.9
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/pkg/errors v0.8.1
	github.com/rs/zerolog v1.15.0
	github.com/sirkon/goproxy v1.4.8
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	github.com/tendermint/tendermint v0.31.10
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.2.8
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
