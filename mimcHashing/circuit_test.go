package mimcHashing

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	"github.com/stretchr/testify/assert"

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
	byteData := data.Bytes()
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

func TestPublicWitness(t *testing.T) {
	dummyHashedData := utils.GetDummyHashedData()
	assignment := MimcCircuit{
		PreImage: dummyHashedData.Data,
		Hash:     utils.HexStringToByteArray(dummyHashedData.HashString),
	}

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())

	publicWitness, _ := witness.Public()

	field := ecc.BN254.ScalarField()

	newWitness, err := assignment.ConstructWitness(field)

	assert.NoError(t, err)

	assert.Equal(t, publicWitness, newWitness)
}
