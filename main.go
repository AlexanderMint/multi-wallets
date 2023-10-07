package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"log"
	"strings"

	"./btc"
	"./eth"
)

// BIP39 Mnemonic: fruit hip zone ridge birth joy reveal venue sword blush edge spice bar subject begin target oval vicious wink village enhance bomb slab cake
const (
	BTCPublicKey32  = "xpub67zki1MNery24HpezRXGxuvpJiuRJYW9WAwCG3gD2ky7PHQhjmksBDrv2zsucMdFo2vT61ngDXE8faEkZmYCBnknK1KALyjJGSsmB9vZHPX"
	BTCPublicKey44  = "xpub6EsDZSyqmejvVo2W7BqfEucLkdpePbkNJvnsdZ4E94TTv5D1EU5ZVx4pB5rrabXABearF2DKyjRMrSkUAHELLhPEVKruJu9xQs92Lj42cqJ"
	BTCPublicKey49  = "ypub6Zek6Qs3WakHqwBRKw7atwdXYmcNVDEZGrQc7gmxGHGkGRskWo58PvR7w3ofCHFE8vVS8KSNYx9m2EUCGE6y1VpvUw6r9zCBoaShYr94hAN"
	BTCPublicKey84  = "zpub6sFUjLu7crB5RmambHx3DdKE81xzxkmW1bBYtcijj6FusCbF8tGJvoFWmCiZokqkibFSz8g3bTSz4BHgzEYcf3THU6tj8GPPGaKb1FEejUJ"
	BTCPublickey141 = "ypub6Sq21g2HoYWVub1mpnJuB12KUh3sFAVeRHTR3Sa6QmLzSPDvzRvRoHX44CqVcGHBCg3FqVPEgBagYrrKHTxCz2SPBM1avtYnYAwQZfVKr3m"
)

func main() {
	publicKey := strings.TrimSpace(BTCPublicKey44)
	publicPrefix := publicKey[0:4]
	key := newKey(publicKey)

	for i := 0; i < 1; i++ {
		accountKey := newAccount(*key, uint32(i))
		accountKey.Version, _ = hex.DecodeString("0488B21E")

		fmt.Printf("[%v] Public Key:\t%v\n", i, hex.EncodeToString(accountKey.Key))

		ethAddress := eth.KeyToEthAddress(*accountKey)
		fmt.Printf("\t[%v][ETH]\t\tAddress: \t%v\n", i, ethAddress)

		if publicPrefix == "xpub" {
			btcAddressP2PKH := btc.KeyToBTCAddressP2PKH(*accountKey) // 1AuguavtKwYsxhLyGtnqWLTYC4JQARrEKa
			fmt.Printf("\t[%v][BTC][P2PKH]\t\tAddress:\t%v\n", i, btcAddressP2PKH)
		}

		btcAddressP2SH := btc.KeyToBTCAddressP2SH(*accountKey)
		fmt.Printf("\t[%v][BTC][P2SH]\t\tAddress:\t%v\n", i, btcAddressP2SH)

		btcAddressP2WPKH := btc.KeyToBTCAddressP2WPKH(*accountKey)
		fmt.Printf("\t[%v][BTC][P2WPKH]\tAddress:\t%v\n", i, btcAddressP2WPKH)
	}
}

func newKey(publicKey string) *bip32.Key {
	masterPublic, err := bip32.B58Deserialize(publicKey)

	if err != nil { log.Fatal(err) }

	return masterPublic
}

func newAccount(key bip32.Key, index uint32) *bip32.Key {
	newKey, err := key.NewChildKey(index)

	if err != nil { log.Fatal(err) }

	return newKey
}
