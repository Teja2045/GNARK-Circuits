package sha256

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/std/math/uints"
	"github.com/consensys/gnark/test"
)

func TestSha256Circuit(t *testing.T) {
	data1 := []byte("good data very good data")
	data2 := []byte("good date very very date")
	digest := sha256.Sum256(data1)

	fmt.Println("digest", len(digest), len(data2))
	validAssignment := Sha256Circuit{
		In:       uints.NewU8Array(data1),
		Expected: uints.NewU8Array(digest[:]),
	}
	invalidAssignment := Sha256Circuit{
		In:       uints.NewU8Array(data2),
		Expected: uints.NewU8Array(digest[:]),
	}

	circuit := &Sha256Circuit{
		In:       make([]uints.U8, len(data2)),
		Expected: make([]uints.U8, len(digest)),
	}

	err := test.IsSolved(
		circuit, &validAssignment, ecc.BN254.ScalarField(),
	)
	if err != nil {
		t.Fatal(err)
	}

	assert := test.NewAssert(t)
	assert.ProverFailed(
		circuit, &invalidAssignment,
	)
}
