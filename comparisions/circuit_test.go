package comparisions

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
)

func TestComparisions(t *testing.T) {
	assert := test.NewAssert(t)
	var circuit ComparisionCircuit

	assert.CheckCircuit(&circuit, test.WithCurves(ecc.BN254))
	assert.ProverSucceeded(&circuit, &ComparisionCircuit{Add: 2}, test.WithCurves(ecc.BN254))
}
