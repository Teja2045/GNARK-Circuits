package mimcHashing

import (
	"github.com/consensys/gnark-crypto/hash"
)

func Hash(data []byte, hashFunc hash.Hash) []byte {
	mimc := hashFunc.New()
	mimc.Write(data)

	return mimc.Sum(nil)
}
