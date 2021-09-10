package main

import (
	"encoding/json"
	"os"

	nearclient "github.com/eteu-technologies/near-api-go/pkg/client"
	"github.com/urfave/cli/v2"
)

var genesis = &cli.Command{
	Name:  "genesis",
	Usage: "Dumps you the current genesis configuration",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "pretty",
			Aliases: []string{"pretty-print"},
			Usage:   "Whether to pretty print the JSON",
		},
	},
	Action: genesisAction,
}

func genesisAction(cctx *cli.Context) (err error) {
	nodeURL, err := ensureNodeURL(cctx)
	if err != nil {
		return err
	}

	client, err := nearclient.NewClient(nodeURL)
	if err != nil {
		return err
	}

	ctx := cctx.Context
	res, err := client.GenesisConfig(ctx)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	if cctx.Bool("pretty") {
		encoder.SetIndent("", "    ")
	}

	err = encoder.Encode(res)

	return
}
