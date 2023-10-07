package eth

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
)

func KeyToEthAddress(key bip32.Key) string {
	publicKey, _ := crypto.DecompressPubkey(key.Key)

	return crypto.PubkeyToAddress(*publicKey).String()
}
