package signature

import (
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/signature/eddsa"
)

func Prover(pk groth16.ProvingKey, pbKey eddsa.PublicKey, data []byte, signature eddsa.Signature) (groth16.Proof, error) {
	startTime := time.Now()
	defer func(t time.Time) {
		elapsed := time.Since(t).Milliseconds()
		println("Time taken to generate proof :", elapsed, "MilliSeconds")
	}(startTime)

	var circuit SignatureCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	assignment := SignatureCircuit{
		PubKey:    pbKey,
		Data:      data,
		Signature: signature,
	}

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, err
	}
	return groth16.Prove(ccs, pk, witness)

}
