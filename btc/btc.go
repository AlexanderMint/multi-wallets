package btc

import (
	"crypto/sha256"
	"github.com/FactomProject/basen"
	"golang.org/x/crypto/ripemd160"
	"io"
)

const (
	OP_0 = 0x00
)
var (
	BitcoinBase58Encoding = basen.NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

//func KeyToBtcAddress(key bip32.Key) string {
//	crypto.DecompressPubkey(key.Key)
//	//fmt.Println("16JrGhLx5bcBSA34kew9V6Mufa4aXhFe9X - ", toAddressP2PKH(key))
//	fmt.Println("34N3tf5m5rdNhW5zpTXNEJucHviFEa8KEq - ", toAddressP2SH())
//	//fmt.Println("34N3tf5m5rdNhW5zpTXNEJucHviFEa8KEq - ", toAddressP2WPKH(key))
//
//	return toAddressP2PKH(key)
//}

func hashSha256(data []byte) ([]byte, error) {
	hasher := sha256.New()
	_, err := hasher.Write(data)
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func hashRipeMD160(data []byte) ([]byte, error) {
	hasher := ripemd160.New()
	_, err := io.WriteString(hasher, string(data))
	if err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

