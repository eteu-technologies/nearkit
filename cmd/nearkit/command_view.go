package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

var viewCall = &cli.Command{
	Name:  "view",
	Usage: "Calls smart contract's view-only function",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "target-id",
			Usage:    "Account to make this call to",
			Required: true,
		},
	},
	Action: viewCallAction,
}

func viewCallAction(cctx *cli.Context) (err error) {
	return errors.New("not implemented")
}
