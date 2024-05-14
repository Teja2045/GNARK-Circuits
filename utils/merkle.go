package utils

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
)

func GetMerkleProofFromBytes(rootBytes []byte, proofBytes [][]byte) merkle.MerkleProof {
	var merkleProof merkle.MerkleProof
	merkleProof.RootHash = rootBytes
	merkleProof.Path = make([]frontend.Variable, len(proofBytes))
	for i := 0; i < len(proofBytes); i++ {
		merkleProof.Path[i] = proofBytes[i]
	}
	return merkleProof
}
