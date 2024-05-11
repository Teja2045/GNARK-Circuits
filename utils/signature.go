package utils

import (
	"hash"
	"log"
	"math/rand"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
)

func Sign(msg []byte, privateKey eddsa.PrivateKey, hFunc hash.Hash) eddsa.Signature {
	signBytes, err := privateKey.Sign(msg, hFunc)
	if err != nil {
		log.Fatal(err)
	}

	var signature eddsa.Signature
	signature.SetBytes(signBytes)
	return signature
}

func Verify(msg []byte, pubKey eddsa.PublicKey, signature []byte, hFunc hash.Hash) (bool, error) {
	return pubKey.Verify(signature, msg, hFunc)
}

func GenerateKeys(i int64) (eddsa.PrivateKey, eddsa.PublicKey) {
	randSrc := rand.NewSource(i)
	r := rand.New(randSrc)

	privatekey, err := eddsa.GenerateKey(r)
	handleErr(err)

	return *privatekey, (*privatekey).PublicKey
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
