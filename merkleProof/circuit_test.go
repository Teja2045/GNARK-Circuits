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

	var circuit MerkleProofCircuit
	circuit.MerkleProof.Path = make([]frontend.Variable, 6)

	hFunc := mimc.NewMiMC()
	var buf bytes.Buffer

	dataSegments := 3
	proofIndex := 0

	for i := byte(1); i <= byte(dataSegments); i++ {
		data := []byte{i}
		hFunc.Reset()
		hFunc.Write(data)
		hash := hFunc.Sum(nil)
		_, err := buf.Write(hash)
		assert.NoError(err)
	}

	fmt.Println("buffer length: ", buf.Len(), "hFunc", hFunc.Size())

	root, proof, numLeaves, err := merkletree.BuildReaderProof(&buf, hFunc, dataSegments, uint64(proofIndex))
	assert.NoError(err)

	// index should be 0
	verified := merkletree.VerifyProof(hFunc, root, proof, 1, numLeaves)
	assert.Equal(verified, false)
	verified = merkletree.VerifyProof(hFunc, root, proof, 0, numLeaves)
	assert.Equal(verified, true)

	assingment := MerkleProofCircuit{
		MerkleProof: utils.GetMerkleProofFromBytes(root, proof),
		Data:        uint64(proofIndex),
	}

	assert.CheckCircuit(&circuit, test.WithValidAssignment(&assingment), test.WithCurves(ecc.BN254))

	// root []byte, proofSet [][]byte, numLeaves uint64, err error

}
