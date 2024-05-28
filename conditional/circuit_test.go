package conditional

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
)

func TestConditionalCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var circuit ConditionalCircuit

	assert.CheckCircuit(&circuit, test.WithCurves(ecc.BN254))

	// try a failed case: sum = 0 + 9 < 10
	assert.ProverSucceeded(&circuit, &ConditionalCircuit{A: 5, B: 10, A_Support: 0, B_Support: 1}, test.WithCurves(ecc.BN254))
}
