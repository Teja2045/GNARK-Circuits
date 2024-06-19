package merkleproof

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Teja2045/GNARK-Circuits/utils"
	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

func TestMerkleProof(t *testing.T) {
	assert := test.NewAssert(t)

	hFunc := mimc.NewMiMC()
	var buf bytes.Buffer

	dataSegments := 4
	proofIndex := 0
	dataSize := 32 // hash size

	for i := byte(1); i <= byte(dataSegments); i++ {
		data := []byte{i}
		hFunc.Reset()
		hFunc.Write(data)
		hash := hFunc.Sum(nil)
		fmt.Println("hash at index", i, hash)

		_, err := buf.Write(hash)
		assert.NoError(err)
	}

	fmt.Println("buffer length: ", buf.Len(), "hFunc", hFunc.Size())

	root, proof, numLeaves, err := merkletree.BuildReaderProof(&buf, hFunc, dataSize, uint64(proofIndex))
	assert.NoError(err)

	fmt.Println("proofLength", len(proof))
	fmt.Println("leaf", proof[0])

	// index should be 0
	verified := merkletree.VerifyProof(hFunc, root, proof, 1, numLeaves)
	assert.Equal(verified, false)
	verified = merkletree.VerifyProof(hFunc, root, proof, 0, numLeaves)
	assert.Equal(verified, true)

	assingment := MerkleProofCircuit{
		MerkleProof: utils.GetMerkleProofFromBytes(root, proof),
		LeafIndex:   uint64(proofIndex),
		LeafData:    proof[0],
	}

	var circuit MerkleProofCircuit
	circuit.MerkleProof.Path = make([]frontend.Variable, len(proof))

	assert.CheckCircuit(&circuit, test.WithValidAssignment(&assingment), test.WithCurves(ecc.BN254))

	//panic("bruh")
	// root []byte, proofSet [][]byte, numLeaves uint64, err error

}

func TestMerkleProofWithSha256(t *testing.T) {

}
