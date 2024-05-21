package recursion

import "github.com/consensys/gnark/frontend"

type InnerCircuit struct {
	P frontend.Variable
	Q frontend.Variable
	R frontend.Variable `gnark:",public"`
}

func (circuit *InnerCircuit) Define(api frontend.API) error {
	res := api.Mul(circuit.P, circuit.Q)
	api.AssertIsEqual(res, circuit.R)
	api.AssertIsDifferent(circuit.P, 1)
	api.AssertIsDifferent(circuit.Q, 1)
	return nil
}
