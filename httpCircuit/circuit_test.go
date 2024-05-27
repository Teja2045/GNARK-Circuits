package httpCircuit

import (
	"testing"

	"github.com/consensys/gnark/test"
)

func TestHttpCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var circuit HttpCircuit

	assert.CheckCircuit(&circuit)
	assert.ProverSucceeded(&circuit, &HttpCircuit{Score: 100000})
}
