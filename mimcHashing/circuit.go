package mimcHashing

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

type MimcCircuit struct {
	PreImage frontend.Variable
	Hash     frontend.Variable `gnark:",public"`
}

func (circuit *MimcCircuit) Define(api frontend.API) error {

	mimc, _ := mimc.NewMiMC(api)

	mimc.Write(circuit.PreImage)

	api.AssertIsEqual(circuit.Hash, mimc.Sum())
	return nil
}
