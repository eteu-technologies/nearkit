# NEARKit

CLI to interact with NEAR blockchain

## Running

### Using Nix

We're using Nix flakes!

`nix run github:eteu-technologies/nearkit -- --help`

### Build & run

Clone and do `go build ./cmd/nearkit`. You'll get binary named `nearkit` what you can install somewhere.

<!--
NOTE: Broken as of 2021-07-25:
go/src/github.com/eteu-technologies/near-api-go/pkg/types/balance.go:9:2: code in directory /Users/mark/go/src/github.com/eteu-technologies/golang-uint128 expects import "lukechampine.com/uint128"


### Using Go CLI

```
GO111MODULE=off go get github.com/eteu-technologies/nearkit/cmd/nearkit
GO111MODULE=off go run github.com/eteu-technologies/nearkit/cmd/nearkit
```
-->

## License

MIT
