package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	nearconfig "github.com/eteu-technologies/near-api-go/pkg/config"
	"github.com/eteu-technologies/nearkit/internal/account"
	"github.com/urfave/cli/v2"
)

func ensureAccountID(cctx *cli.Context) (accountID string, err error) {
	accountID = cctx.String("account-id")
	if accountID == "" {
		err = errors.New("account ID is not set, set it using --account-id flag")
	}
	return
}

func ensureCredential(cctx *cli.Context) (*account.Credential, error) {
	accountID, err := ensureAccountID(cctx)
	if err != nil {
		return nil, err
	}

	credentialsFile := cctx.Path("credentials-file")

	// Try resolving the credentials file
	if credentialsFile == "" {
		networkID := cctx.String("network-id")

		// Ensure credentials directory is set
		var credentialsDir string
		if credentialsDir = cctx.Path("credentials-directory"); credentialsDir == "" {
			if networkID == "" {
				return nil, errors.New("Using custom network RPC URL, cannot determine network ID and credentials path reliably. Please specify it using '--credentials-directory'")
			}

			credentialsDir = cctx.String("credentials-base-directory")
			if credentialsDir == "" {
				credentialsDir, err = os.UserHomeDir()
				if err != nil {
					return nil, fmt.Errorf("Unable to resolve home directory, try using '--credentials-base-directory'? %w", err)
				}
			}

			credentialsDir = filepath.Join(credentialsDir, ".near-credentials", networkID)
			_ = cctx.Set("credentials-directory", credentialsDir)
		}

		// Resolve the file
		credentialsFile = filepath.Join(credentialsDir, fmt.Sprintf("%s.json", accountID))
		_ = cctx.Set("credentials-file", credentialsFile)
	}

	credential, err := account.LoadCredentials(credentialsFile)
	if err != nil {
		return nil, err
	}

	return &credential, err
}

func ensureNodeURL(cctx *cli.Context) (nodeURL string, err error) {
	networkID := cctx.String("network-id")

	// Ensure network-rpc-url is set
	if !cctx.IsSet("network-rpc-url") {
		// Check if network is valid
		network, ok := nearconfig.Networks[networkID]
		if !ok {
			return nodeURL, fmt.Errorf("Invalid network %s", networkID)
		}

		nodeURL = network.NodeURL
		log.Printf("using network '%s'; rpc url '%s'", network.NetworkID, network.NodeURL)
		_ = cctx.Set("network-rpc-url", network.NodeURL)
	} else {
		// Unset networkID
		_ = cctx.Set("network-id", "")
		nodeURL = cctx.String("network-rpc-url")
	}
	return
}
