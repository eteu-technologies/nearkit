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
		PrivateKey string `json:"private_key"`
		PublicKey  string `json:"public_key"`
	}

	keyPair, err := nearkey.GenerateKeyPair(nearkey.KeyTypeED25519, rand.Reader)
	if err != nil {
		err = fmt.Errorf("Failed to generate ed25519 keypair: %w", err)
		return
	}

	key.AccountID = cctx.String("account")
	key.PrivateKey = keyPair.PrivateEncoded()
	key.PublicKey = keyPair.PublicKey.String()

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	if cctx.Bool("pretty") {
		encoder.SetIndent("", "    ")
	}
	err = encoder.Encode(&key)

	return
}
