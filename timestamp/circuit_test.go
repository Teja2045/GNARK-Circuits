package timestamp

import (
	"testing"
	"time"

	"github.com/consensys/gnark/test"
)

func TestTimeCirucit(t *testing.T) {
	assert := test.NewAssert(t)

	curTime := time.Now()
	curTimeStamp := curTime.Unix()

	assignment := TimeCircuit{TimeStamp: curTimeStamp}
	<-time.After(time.Second)
	assert.CheckCircuit(&TimeCircuit{})
	assert.ProverSucceeded(&TimeCircuit{}, &assignment)

	// panic("")
}
