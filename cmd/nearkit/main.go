package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "nearkit",
		Usage: "Interact with NEAR blockchain",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "network-id",
				Aliases: []string{"network"},
				EnvVars: []string{
					"NEARKIT_NETWORK_ID",
					"NEAR_ENV", // near-cli compatibility
				},
				Value: "testnet",
				Usage: "NEAR network to use",
			},
			&cli.StringFlag{
				Name: "network-rpc-url",
				EnvVars: []string{
					"NEARKIT_NETWORK_RPC_URL",
				},
				Usage: "NEAR network RPC URL",
			},
			/*
				&cli.StringFlag{
					Name:  "output-format",
					Usage: "Output format to use",
					Value: "human",
				},
			*/
			&cli.PathFlag{
				Name: "credentials-directory",
				EnvVars: []string{
					"NEARKIT_CREDENTIALS_DIRECTORY",
				},
				Usage: "Path to where credentials are stored",
			},
			&cli.PathFlag{
				Name: "credentials-base-directory",
				EnvVars: []string{
					"NEARKIT_CREDENTIALS_BASE_DIRECTORY",
				},
				Usage: "Path to substitute as home directory (~/.near-credentials)",
			},
			&cli.PathFlag{
				Name: "credentials-file",
				EnvVars: []string{
					"NEARKIT_CREDENTIALS_FILE",
				},
				Usage:       "Path to the credentials file for current action",
				DefaultText: "Will be resolved according to the network id",
			},
			&cli.StringFlag{
				Name: "account-id",
				EnvVars: []string{
					"NEARKIT_ACCOUNT_ID",
				},
				Aliases: []string{"accountId"}, // near-cli compatibility
				Usage:   "Account ID to use for transaction signing",
			},
		},
		Commands: []*cli.Command{
			deployContract,
			viewCall,
			changeCall,
			genesis,
		},
		Before: before,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func before(cctx *cli.Context) (err error) {
	return
}
