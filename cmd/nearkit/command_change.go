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
			Name:     "gas",
			Usage:    "Amount of gas to attach",
			Value:    neartypes.DefaultFunctionCallGas,
			Required: true,
		},
		&cli.StringFlag{
			Name:  "deposit",
			Usage: "Amount of NEAR tokens to attach",
			Value: "0",
		},
		&cli.StringFlag{
			Name:     "target-id",
			Usage:    "Account to make this call to",
			Required: true,
		},
	},
	Action: changeCallAction,
}

func changeCallAction(cctx *cli.Context) (err error) {
	return errors.New("not implemented")
}
