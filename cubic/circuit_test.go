package cubic

import (
	"testing"

	"github.com/consensys/gnark/test"
)

func TestCubicCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var cubicCircuit CubicCircuit

	assert.ProverSucceeded(&cubicCircuit, &CubicCircuit{
		X: 1,
		Y: 1,
	})

	assert.ProverSucceeded(&cubicCircuit, &CubicCircuit{
		X: 2,
		Y: 8,
	})

	assert.ProverFailed(&cubicCircuit, &CubicCircuit{
		X: 2,
		Y: 9,
	})
}
