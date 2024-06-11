package secp256k1_ecdsa

import (
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/secp256k1/ecdsa"
	"github.com/consensys/gnark/std/math/emulated"
	gnark_ecdsa "github.com/consensys/gnark/std/signature/ecdsa"
	"github.com/consensys/gnark/test"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func TestSecp256k1Circuit(t *testing.T) {
	secret := []byte("this is a huge secret")
	sdkPrivkey := secp256k1.GenPrivKeyFromSecret(secret)
	sdkPubkey := sdkPrivkey.PubKey()

	// this method exists in forked gnark-crypto
	gnarkPrivkey, _ := ecdsa.GenerateKeyFromScalar(sdkPrivkey.Key)
	gnarkPubkey := gnarkPrivkey.PublicKey

	// msg to sign
	msg := []byte("message")
	hFunc := sha256.New()
	hFunc.Reset()

	sign, err := sdkPrivkey.Sign(msg)
	if err != nil {
		t.Errorf("some thing went wrong %t", err)
	}

	flag := sdkPubkey.VerifySignature(msg, sign)
	if !flag {
		t.Errorf("can't verify signature using sdk public key")
	}

	flag, _ = gnarkPubkey.Verify(sign, msg, hFunc)
	if !flag {
		t.Errorf("can't verify signature using gnark public key")
	}

	var sig ecdsa.Signature
	sig.SetBytes(sign)
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(sig.R[:32])
	s.SetBytes(sig.S[:32])

	dataToHash := make([]byte, len(msg))
	copy(dataToHash, msg)
	hFunc.Reset()
	hFunc.Write(dataToHash)
	hramBin := hFunc.Sum(nil)
	hash := ecdsa.HashToInt(hramBin)

	circuit := Secp256k1Circuit[emulated.Secp256k1Fp, emulated.Secp256k1Fr]{}
	witness := Secp256k1Circuit[emulated.Secp256k1Fp, emulated.Secp256k1Fr]{
		Signature: gnark_ecdsa.Signature[emulated.Secp256k1Fr]{
			R: emulated.ValueOf[emulated.Secp256k1Fr](r),
			S: emulated.ValueOf[emulated.Secp256k1Fr](s),
		},
		Message: emulated.ValueOf[emulated.Secp256k1Fr](hash),
		Pubkey: gnark_ecdsa.PublicKey[emulated.Secp256k1Fp, emulated.Secp256k1Fr]{
			X: emulated.ValueOf[emulated.Secp256k1Fp](gnarkPubkey.A.X),
			Y: emulated.ValueOf[emulated.Secp256k1Fp](gnarkPubkey.A.Y),
		},
	}

	assert := test.NewAssert(t)
	err = test.IsSolved(&circuit, &witness, ecc.BN254.ScalarField())
	assert.ProverSucceeded(&circuit, &witness)
	assert.NoError(err)

}

// https://github.com/Teja2045/gnark-crypto/tree/master/ecc/secp256k1/ecdsa
