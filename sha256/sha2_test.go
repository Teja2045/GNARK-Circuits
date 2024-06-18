package sha256

import (
	"crypto/sha256"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/std/math/uints"
	"github.com/consensys/gnark/test"
)

func TestSha256Circuit(t *testing.T) {
	data1 := []byte("good data very good data")
	data2 := []byte("goo  data vere good date")
	digest := sha256.Sum256(data1)
	assignment := Sha256Circuit{
		In: uints.NewU8Array(data2),
	}

	copy(assignment.Expected, uints.NewU8Array(digest[:]))

	err := test.IsSolved(&Sha256Circuit{In: make([]uints.U8, len(data2))}, &assignment, ecc.BN254.ScalarField())
	if err != nil {
		t.Fatal(err)
	}

	assert := test.NewAssert(t)
	assert.ProverFailed(&Sha256Circuit{In: make([]uints.U8, len(data1))}, &assignment)
}
