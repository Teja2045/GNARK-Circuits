package signature

import (
	"log"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/signature/eddsa"
)

func Verifier(proof groth16.Proof, vk groth16.VerifyingKey, data []byte, randomPubKey eddsa.PublicKey, randomSign eddsa.Signature) {

	startTime := time.Now()
	defer func(t time.Time) {
		elapsed := time.Since(t).Milliseconds()
		println("Time taken to verify proof:", elapsed, " MilliSeconds")
	}(startTime)

	// make an assignment with valid public inputs and other random inputs
	publicAssignment := SignatureCircuit{
		PubKey:    randomPubKey,
		Signature: randomSign,
		Data:      data,
	}

	witness, err := frontend.NewWitness(&publicAssignment, ecc.BN254.ScalarField())
	if err != nil {
		log.Println(err)
	}
	publicWitness, _ := witness.Public()

	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		log.Fatal(err)
	}
}
