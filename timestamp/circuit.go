package timestamp

import (
	"time"

	"github.com/consensys/gnark/frontend"
)

type TimeCircuit struct {
	TimeStamp frontend.Variable
}

// this is wrong circuit as time.Now is dynamic, it should be constant inside the circuit
// change time.Unix to time.UnixNano to observe the error
func (circuit *TimeCircuit) Define(api frontend.API) error {
	now := time.Now()
	timeStampNow := now.Unix()
	api.AssertIsLessOrEqual(circuit.TimeStamp, timeStampNow)

	// wrong assertion
	// api.AssertIsLessOrEqual(timeStampNow, circuit.TimeStamp)

	api.AssertIsDifferent(timeStampNow, circuit.TimeStamp)
	return nil
}
