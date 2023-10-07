package btc

import (
	"github.com/tyler-smith/go-bip32"
)

// KeyToBTCAddressP2SH (Pay To Script Hash)
// https://ru.bitcoinwiki.org/wiki/P2SH
func KeyToBTCAddressP2SH(key bip32.Key) string {
	//pubBytes, _ := hex.DecodeString("3a38d44d6a0c8d0bb84e0232cc632b7e48c72e0e")
	//keyhash := append([]byte{OP_0, byte(len(pubBytes))}, pubBytes...) // 00143a38d44d6a0c8d0bb84e0232cc632b7e48c72e0e
	keyhash, _ := hashSha256(key.Key)  // 1ae968057eaef06c3e13439695edd7a54982fc99f36c3aa41d8cc41340f30195
	keyhash, _ = hashRipeMD160(keyhash) // 1d521dcf4983772b3c1e6ef937103ebdfaa1ad77

	addrhash := append([]byte{0x05}, keyhash...) // 051d521dcf4983772b3c1e6ef937103ebdfaa1ad77

	checksum, _ := hashSha256(addrhash)
	checksum, _ = hashSha256(checksum)

	address := append(addrhash, checksum[:4]...)

	return BitcoinBase58Encoding.EncodeToString(address) // 34N3tf5m5rdNhW5zpTXNEJucHviFEa8KEq
}
