package account

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadCredentials(credentialFile string) (credential Credential, err error) {
	var cf *os.File
	if cf, err = os.Open(credentialFile); err != nil {
		return
	}
	defer cf.Close()

	if err = json.NewDecoder(cf).Decode(&credential); err != nil {
		return
	}

	if credential.PublicKey.String() != credential.PrivateKey.PublicKey.String() {
		err = fmt.Errorf("inconsistent public key, %s != %s", credential.PublicKey.String(), credential.PrivateKey.PublicKey.String())
		return
	}

	return
}
