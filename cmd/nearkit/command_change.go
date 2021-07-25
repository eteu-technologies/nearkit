package main

import (
	"errors"

	neartypes "github.com/eteu-technologies/near-api-go/pkg/types"
	"github.com/urfave/cli/v2"
)

var changeCall = &cli.Command{
	Name:    "change",
	Aliases: []string{"call"},
	Usage:   "Calls smart contract's function",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name: "gas",
			EnvVars: []string{
				"NEARKIT_CALL_ATTACH_GAS",
			},
			Usage:    "Amount of gas to attach",
			Value:    neartypes.DefaultFunctionCallGas,
			Required: true,
		},
		&cli.StringFlag{
			Name: "deposit",
			EnvVars: []string{
				"NEARKIT_CALL_ATTACH_DEPOSIT",
			},
			Usage: "Amount of NEAR tokens to attach",
			Value: "0",
		},
		&cli.StringFlag{
			Name: "target-id",
			EnvVars: []string{
				"NEARKIT_TARGET_ACCOUNT_ID",
			},
			Usage:    "Account to make this call to",
			Required: true,
		},
	},
	Action: changeCallAction,
}

func changeCallAction(cctx *cli.Context) (err error) {
	return errors.New("not implemented")
}
