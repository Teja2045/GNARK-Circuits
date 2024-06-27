package utils

import (
	"fmt"

	"github.com/consensys/gnark/backend/witness"
)

func WitnessToInt(pbWitness witness.Witness) uint32 {
	data, _ := pbWitness.MarshalBinary()
	fmt.Println(data)
	fmt.Println(len(data))
	return 0
}
