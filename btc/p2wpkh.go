package btc

import (
	"encoding/hex"
	"github.com/tyler-smith/go-bip32"
)

// KeyToBTCAddressP2WPKH (Pay To Public Key Hash)
// For test https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses
// https://ru.bitcoinwiki.org/wiki/P2PKH
func KeyToBTCAddressP2WPKH(key bip32.Key) string {
	pubBytes, _ := hex.DecodeString("02530c548d402670b13ad8887ff99c294e67fc18097d236d57880c69261b42def7")
	keyhash := append([]byte{OP_0, byte(len(pubBytes))}, pubBytes...)
	keyhash, _ = hashSha256(key.Key)
	keyhash, _ = hashRipeMD160(keyhash)

	addrhash := append([]byte{0x00, 0x14}, keyhash...)

	checksum, _ := hashSha256(addrhash)
	checksum, _ = hashSha256(checksum)

	addr := append(addrhash, checksum[:4]...)

	return BitcoinBase58Encoding.EncodeToString(addr) // bc1qg9stkxrszkdqsuj92lm4c7akvk36zvhqw7p6ck
}