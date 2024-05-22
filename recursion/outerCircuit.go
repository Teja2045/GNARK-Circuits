package recursion

import (
	"fmt"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra"
	"github.com/consensys/gnark/std/math/emulated"
	stdgroth16 "github.com/consensys/gnark/std/recursion/groth16"
)

type OuterCircuit[
	FR emulated.FieldParams,
	G1El algebra.G1ElementT,
	G2El algebra.G2ElementT,
	GtEl algebra.GtElementT] struct {
	Proof        stdgroth16.Proof[G1El, G2El]
	VerifyingKey stdgroth16.VerifyingKey[G1El, G2El, GtEl]
	InnerWitness stdgroth16.Witness[FR] `gnark:",public"`
	P            frontend.Variable
	Q            frontend.Variable
	R            frontend.Variable `gnark:",public"`
}

func (circuit *OuterCircuit[FR, G1El, G2El, GtEl]) Define(api frontend.API) error {
	verifier, err := stdgroth16.NewVerifier[FR, G1El, G2El, GtEl](api)
	if err != nil {
		return fmt.Errorf("new Verifier: %w", err)
	}
	circuit.primeCheck(api)
	return verifier.AssertProof(circuit.VerifyingKey, circuit.Proof, circuit.InnerWitness)
}

func (circuit *OuterCircuit[FR, G1El, G2El, GtEl]) primeCheck(api frontend.API) {
	res := api.Mul(circuit.P, circuit.Q)
	api.AssertIsEqual(res, circuit.R)
	api.AssertIsDifferent(circuit.P, 1)
	api.AssertIsDifferent(circuit.Q, 1)
}
