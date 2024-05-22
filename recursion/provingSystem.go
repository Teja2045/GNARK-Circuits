package recursion

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	stdgroth16 "github.com/consensys/gnark/std/recursion/groth16"
)

func ProveAndVerify() {
	field := ecc.BN254.ScalarField()
	innerPk, innerVk, innerCs, err := PersistantInnerSetup()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	innerAssignment := NewInnerAssignment(3, 5, 15)
	innerPubWitness, innerProof := GenerateProof(field, innerPk, innerCs, &innerAssignment, true)

	before := time.Now()
	err = groth16.Verify(innerProof, innerVk, innerPubWitness, stdgroth16.GetNativeVerifierOptions(field, field))
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(fmt.Sprintln("time taken taken to verify inner proof: ", time.Since(before).Milliseconds(), "Millseconds"))
	fmt.Println()

	primesData := DummyPrimesData()

	var outerPubWitness witness.Witness
	var outerProof groth16.Proof

	outerPk, outerVk, outerCs, err := PersistantOuterSetup(innerCs)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for i := 0; i < len(primesData); i++ {
		outerAssignment := NewOuterAssignment(primesData[i][0], primesData[i][1], primesData[i][2], innerPubWitness, innerProof, innerVk)

		outerPubWitness, outerProof = GenerateProof(field, outerPk, outerCs, &outerAssignment, i+1 < len(primesData))
		innerPubWitness = outerPubWitness
		innerProof = outerProof
		innerVk = outerVk
	}

	before = time.Now()
	groth16.Verify(outerProof, outerVk, outerPubWitness)
	slog.Info(fmt.Sprintln("time taken taken to verify inner proof: ", time.Since(before).Milliseconds(), "Millseconds"))

}
