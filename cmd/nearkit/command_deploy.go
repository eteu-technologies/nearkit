package main

import (
	"io/ioutil"
	"log"

	nearclient "github.com/eteu-technologies/near-api-go/pkg/client"
	nearconfig "github.com/eteu-technologies/near-api-go/pkg/config"
	nearaction "github.com/eteu-technologies/near-api-go/pkg/types/action"
	"github.com/urfave/cli/v2"
)

var deployContract = &cli.Command{
	Name:  "deploy",
	Usage: "Deploys a smart contract under specified account",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name: "wasm-file",
			EnvVars: []string{
				"NEARKIT_DEPLOY_WASM_FILE",
			},
			Aliases:  []string{"wasmFile"}, // near-cli compatibility
			Usage:    "WASM blob to deploy",
			Required: true,
		},
	},
	Action: deployContractAction,
}

func deployContractAction(cctx *cli.Context) (err error) {
	nodeURL, err := ensureNodeURL(cctx)
	if err != nil {
		return err
	}

	credential, err := ensureCredential(cctx)
	if err != nil {
		return err
	}

	wasmFile := cctx.Path("wasm-file")
	wasmBlob, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		return err
	}

	log.Printf("Deploying '%s' to account '%s' (network '%s')", wasmFile, credential.AccountID, nodeURL)

	client, err := nearclient.NewClient(nodeURL)
	if err != nil {
		return err
	}

	ctx := nearclient.ContextWithKeyPair(cctx.Context, credential.PrivateKey)
	res, err := client.TransactionSendAwait(
		ctx,
		credential.AccountID,
		credential.AccountID,
		[]nearaction.Action{
			nearaction.NewDeployContract(wasmBlob),
		},
		nearclient.WithLatestBlock(),
	)
	if err != nil {
		return err
	}

	if networkInfo, ok := nearconfig.Networks[cctx.String("network-id")]; ok {
		log.Printf("%s/transactions/%s", networkInfo.ExplorerURL, res.Transaction.Hash)
	} else {
		log.Printf("%s", res.Transaction.Hash)
	}

	return
}
