package btc

import "github.com/tyler-smith/go-bip32"

const prefix = "1"

// KeyToBTCAddressP2PKH (Pay To Public Key Hash) BIP44
// For test https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses
// https://ru.bitcoinwiki.org/wiki/P2PKH
func KeyToBTCAddressP2PKH(key bip32.Key) string {
	//pubBytes, _ := hex.DecodeString("023cba1f4d12d1ce0bced725373769b2262c6daa97be6a0588cfec8ce1a5f0bd09")
	//keyhash := append([]byte{OP_0, byte(len(pubBytes))}, pubBytes...)
	keyhash, _ := hashSha256(key.Key) // 8eb001a42122826648e66005a549fc4b4511a7ad3fc378221aa1c73c5efe77ef
	keyhash, _ = hashRipeMD160(keyhash) // 3a38d44d6a0c8d0bb84e0232cc632b7e48c72e0e

	addrhash := append([]byte{0x00}, keyhash...) // 003a38d44d6a0c8d0bb84e0232cc632b7e48c72e0e

	checksum, _ := hashSha256(addrhash)
	checksum, _ = hashSha256(checksum)

	addr := append(addrhash, checksum[:4]...)

	return prefix + BitcoinBase58Encoding.EncodeToString(addr) // 16JrGhLx5bcBSA34kew9V6Mufa4aXhFe9X
}