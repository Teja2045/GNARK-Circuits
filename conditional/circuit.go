package conditional

import (
	"github.com/consensys/gnark/frontend"
)

// circuit that does conditional additon and check if the sum
// is greater than 10
type ConditionalCircuit struct {
	A         frontend.Variable
	B         frontend.Variable
	A_Support frontend.Variable
	B_Support frontend.Variable
}

func (circuit *ConditionalCircuit) Define(api frontend.API) error {

	// equivalent to circuit.A_support == 1 ? circuit.A : 0
	toAdd1 := api.Select(circuit.A_Support, circuit.A, 0)
	toAdd2 := api.Select(circuit.B_Support, circuit.B, 0)

	sum := api.Add(toAdd1, toAdd2)
	api.AssertIsLessOrEqual(10, sum)
	return nil
}
