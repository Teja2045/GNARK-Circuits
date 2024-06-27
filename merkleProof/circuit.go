package merkleproof

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
)

type MerkleProofCircuit struct {
	LeafData    frontend.Variable
	MerkleProof merkle.MerkleProof `gnark:",public"`
	LeafIndex   frontend.Variable  `gnark:",public"`
}

func (circuit *MerkleProofCircuit) Define(api frontend.API) error {

	// hFunc, err := mimc.NewMiMC(api)
	// if err != nil {
	// 	return err
	// }

	api.AssertIsEqual(circuit.LeafData, circuit.MerkleProof.Path[0])

	//circuit.MerkleProof.VerifyProof(api, &hFunc, circuit.LeafIndex)
	return nil
}
