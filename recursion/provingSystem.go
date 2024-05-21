package recursion

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	stdgroth16 "github.com/consensys/gnark/std/recursion/groth16"
)

func ProveAndVerify() {
	field := ecc.BN254.ScalarField()
	innerPk, innerVk, innerCs, err := InnerSetup()
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

	outerAssignment := NewOuterAssignment(innerPubWitness, innerProof, innerVk)
	outerPk, outerVk, outerCs, err := OuterSetup(innerCs)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	outerPubWitness, outerProof := GenerateProof(field, outerPk, outerCs, &outerAssignment, false)

	before = time.Now()
	groth16.Verify(outerProof, outerVk, outerPubWitness)
	slog.Info(fmt.Sprintln("time taken taken to verify inner proof: ", time.Since(before).Milliseconds(), "Millseconds"))

}
