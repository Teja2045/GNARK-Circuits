package merkleproof

import (
	"github.com/Teja2045/GNARK-Circuits/sha256"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
)

type Sha256MerkleProofCircuit struct {
	LeafData    frontend.Variable
	MerkleProof merkle.MerkleProof `gnark:",public"`
	LeafIndex   frontend.Variable  `gnark:",public"`
}

func (circuit *Sha256MerkleProofCircuit) Define(api frontend.API) error {

	hFunc, err := sha256.New(api)
	if err != nil {
		return err
	}

	api.AssertIsEqual(circuit.LeafData, circuit.MerkleProof.Path[0])

	circuit.MerkleProof.VerifyProof(api, hFunc, circuit.LeafIndex)
	return nil
}
