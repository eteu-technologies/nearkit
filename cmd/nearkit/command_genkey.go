package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	nearkey "github.com/eteu-technologies/near-api-go/pkg/types/key"
	"github.com/urfave/cli/v2"
)

var genKey = &cli.Command{
	Name:  "genkey",
	Usage: "Generates ED25519 key pair, optionally with account ID",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"pretty-print"},
			Usage:   "Whether to pretty print the JSON",
		},
		&cli.BoolFlag{
			Name:  "node-key",
			Usage: `Whether key is meant to use with NEAR node. Renames "private_key" to "secret_key"`,
		},
		&cli.StringFlag{
			Name:  "account",
			Usage: "Account ID to use (optional)",
		},
	},
	Action: genKeyAction,
}

func genKeyAction(cctx *cli.Context) (err error) {
	var key struct {
		AccountID  string `json:"account_id,omitempty"`
		SecretKey  string `json:"secret_key,omitempty"`
		PrivateKey string `json:"private_key,omitempty"`
		PublicKey  string `json:"public_key"`
	}

	keyPair, err := nearkey.GenerateKeyPair(nearkey.KeyTypeED25519, rand.Reader)
	if err != nil {
		err = fmt.Errorf("Failed to generate ed25519 keypair: %w", err)
		return
	}

	nodeKey := cctx.Bool("node-key")

	key.AccountID = cctx.String("account")
	if key.AccountID == "" && nodeKey {
		key.AccountID = "node"
	}

	if pk := keyPair.PrivateEncoded(); nodeKey {
		key.SecretKey = pk
	} else {
		key.PrivateKey = pk
	}
	key.PublicKey = keyPair.PublicKey.String()

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	if cctx.Bool("pretty") {
		encoder.SetIndent("", "    ")
	}
	err = encoder.Encode(&key)

	return
}
