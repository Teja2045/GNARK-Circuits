package mimcHashing

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/test"

	"zkCircuits/utils"
)

func TestMimcCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var circuit MimcCircuit

	curve := ecc.BN254
	hashFunc := hash.MIMC_BN254

	// hashing using mimc hash and ecc.BN254 curve
	dummyHashedData := utils.GetDummyHashedData()
	data := dummyHashedData.Data
	byteData := []byte(data.Bytes())
	expectedHash := Hash(byteData, hashFunc)

	t.Log("data: ", data, ", expected hash:", utils.ByteArrayToHexString(expectedHash), dummyHashedData.HashString)

	assert.CheckCircuit(&circuit, test.WithValidAssignment(&MimcCircuit{
		PreImage: data,
		Hash:     expectedHash,
	}), test.WithCurves(curve))

	assert.ProverSucceeded(&circuit, &MimcCircuit{
		PreImage: data,
		Hash:     expectedHash,
	}, test.WithCurves(curve))

	assert.ProverFailed(&circuit, &MimcCircuit{
		PreImage: data,
		Hash:     1,
	})

}
