package merkleproof

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/std/hash/mimc"
)

type MerkleProofCircuit struct {
	MerkleProof merkle.MerkleProof `gnark:",public"`
	Data        frontend.Variable  `gnark:",public"`
}

func (circuit *MerkleProofCircuit) Define(api frontend.API) error {

	hFunc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	circuit.MerkleProof.VerifyProof(api, &hFunc, circuit.Data)
	return nil
}
