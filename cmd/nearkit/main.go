package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	nearconfig "github.com/eteu-technologies/near-api-go/pkg/config"
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
				Aliases:  []string{"accountId"}, // near-cli compatibility
				Usage:    "Account ID to use for transaction signing",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			deployContract,
			viewCall,
			changeCall,
		},
		Before: before,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func before(cctx *cli.Context) (err error) {
	networkID := cctx.String("network-id")

	// Ensure network-rpc-url is set
	if !cctx.IsSet("network-rpc-url") {
		// Check if network is valid
		network, ok := nearconfig.Networks[networkID]
		if !ok {
			return fmt.Errorf("Invalid network %s", networkID)
		}

		log.Printf("using network '%s'; rpc url '%s'", network.NetworkID, network.NodeURL)
		if err = cctx.Set("network-rpc-url", network.NodeURL); err != nil {
			return
		}
	} else {
		// Unset networkID
		networkID = ""
		_ = cctx.Set("network", "")
	}

	credentialsFile := cctx.Path("credentials-file")

	// Try resolving the credentials file
	if credentialsFile == "" {
		// Ensure credentials directory is set
		var credentialsDir string
		if credentialsDir = cctx.Path("credentials-directory"); credentialsDir == "" {
			if networkID == "" {
				return errors.New("Using custom network RPC URL, cannot determine network ID and credentials path reliably. Please specify it using '--credentials-directory'")
			}

			credentialsDir = cctx.String("credentials-base-directory")
			if credentialsDir == "" {
				credentialsDir, err = os.UserHomeDir()
				if err != nil {
					return fmt.Errorf("Unable to resolve home directory, try using '--credentials-base-directory'? %w", err)
				}
			}

			credentialsDir = filepath.Join(credentialsDir, ".near-credentials", networkID)

			if err = cctx.Set("credentials-directory", credentialsDir); err != nil {
				return err
			}
		}

		// Resolve the file
		credentialsFile = filepath.Join(credentialsDir, fmt.Sprintf("%s.json", cctx.String("account-id")))
		if err = cctx.Set("credentials-file", credentialsFile); err != nil {
			return
		}
	}

	return
}
