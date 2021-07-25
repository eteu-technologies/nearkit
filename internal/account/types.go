package account

import (
	neartypes "github.com/eteu-technologies/near-api-go/pkg/types"
	nearkey "github.com/eteu-technologies/near-api-go/pkg/types/key"
)

type Credential struct {
	AccountID  neartypes.AccountID     `json:"account_id"`
	PublicKey  nearkey.Base58PublicKey `json:"public_key"`
	PrivateKey nearkey.KeyPair         `json:"private_key"`
}
