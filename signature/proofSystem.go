package signature

import (
	"fmt"
	"log"
	"time"

	"github.com/Teja2045/GNARK-Circuits/utils"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/signature/eddsa"
)

func ProveAndVerify() {
	startTime := time.Now()
	defer func(t time.Time) {
		elapsed := time.Since(t).Milliseconds()
		println("Time taken to for complete Circuit cycle:", elapsed, "MilliSeconds")
	}(startTime)

	var circuit SignatureCircuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		log.Fatal(err)
	}

	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		log.Fatal(err)
	}
	data, pbKey, sign := GetRandomSigntureData()

	proof, err := Prover(pk, pbKey, data, sign)
	fmt.Println("proof generated")

	if err != nil {
		log.Fatal(err)
	}

	_, pbKey2, sign2 := GetRandomSigntureData()
	Verifier(proof, vk, data, pbKey2, sign2)

}

func GetRandomSigntureData() ([]byte, eddsa.PublicKey, eddsa.Signature) {
	mimc := hash.MIMC_BN254
	hFunc := mimc.New()

	data := []byte{1, 2, 3}
	privKey, pubKey := utils.GenerateKeys(0)
	signature := utils.Sign(data, privKey, hFunc)

	var pbKey eddsa.PublicKey
	pbKey.A.X = pubKey.A.X
	pbKey.A.Y = pubKey.A.Y

	var sign eddsa.Signature

	sign.R.X = signature.R.X
	sign.R.Y = signature.R.Y
	sign.S = signature.S[:]

	return data, pbKey, sign
}

func GetRandomSigntureData2() ([]byte, eddsa.PublicKey, eddsa.Signature) {
	mimc := hash.MIMC_BN254
	hFunc := mimc.New()

	data := []byte{1, 2, 3, 4}
	privKey, pubKey := utils.GenerateKeys(1)
	signature := utils.Sign(data, privKey, hFunc)

	var pbKey eddsa.PublicKey
	pbKey.A.X = pubKey.A.X
	pbKey.A.Y = pubKey.A.Y

	var sign eddsa.Signature

	sign.R.X = signature.R.X
	sign.R.Y = signature.R.Y
	sign.S = signature.S[:]

	return data, pbKey, sign
}
