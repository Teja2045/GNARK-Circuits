package recursion

import (
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"time"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	stdgroth16 "github.com/consensys/gnark/std/recursion/groth16"
)

func GenerateProof(
	field *big.Int,
	pk groth16.ProvingKey,
	cs constraint.ConstraintSystem,
	assignment frontend.Circuit,
	isInner bool) (
	witness.Witness,
	groth16.Proof) {

	start := time.Now()
	var proofType string
	if isInner {
		proofType = "inner"
	} else {
		proofType = "outer"
	}
	defer func(before time.Time) {
		timeTaken := time.Since(before).Seconds()
		slog.Info(fmt.Sprintln("time taken to generate", proofType, "proof: ", timeTaken, "seconds"))
	}(start)

	witness, err := frontend.NewWitness(assignment, field)
	if err != nil {
		log.Fatal(fmt.Errorf(err.Error()))
	}

	publicWitness, err := witness.Public()
	if err != nil {
		log.Fatal(fmt.Errorf(err.Error()))
	}
	var proof groth16.Proof
	if isInner {
		proof, err = groth16.Prove(cs, pk, witness, stdgroth16.GetNativeProverOptions(field, field))
	} else {
		proof, err = groth16.Prove(cs, pk, witness)
	}
	if err != nil {
		log.Fatal(fmt.Errorf(err.Error()))
	}

	return publicWitness, proof
}
