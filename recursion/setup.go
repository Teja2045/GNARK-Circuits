package recursion

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/algebra/emulated/sw_bn254"
	stdgroth16 "github.com/consensys/gnark/std/recursion/groth16"
)

func InnerSetup() (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {
	start := time.Now()
	defer func(before time.Time) {
		timeTaken := time.Since(before).Seconds()
		slog.Info(fmt.Sprintln("time taken for inner setup: ", timeTaken, "seconds"))
	}(start)
	var circuit InnerCircuit
	return Setup(&circuit)
}

func OuterSetup(innerCcs constraint.ConstraintSystem) (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {
	start := time.Now()
	defer func(before time.Time) {
		timeTaken := time.Since(before).Seconds()
		slog.Info(fmt.Sprintln("time taken for outer setup: ", timeTaken, "seconds"))
	}(start)
	outerCircuit := &OuterCircuit[sw_bn254.ScalarField, sw_bn254.G1Affine, sw_bn254.G2Affine, sw_bn254.GTEl]{
		InnerWitness: stdgroth16.PlaceholderWitness[sw_bn254.ScalarField](innerCcs),
		VerifyingKey: stdgroth16.PlaceholderVerifyingKey[sw_bn254.G1Affine, sw_bn254.G2Affine, sw_bn254.GTEl](innerCcs),
	}
	return Setup(outerCircuit)
}

func Setup(circuit frontend.Circuit) (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {

	cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("inner proof setup failed: %w", err)
	}

	pk, vk, err := groth16.Setup(cs)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("inner proof setup failed: %w", err)
	}

	return pk, vk, cs, nil

}
