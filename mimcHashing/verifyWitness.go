package mimcHashing

import (
	"math/big"

	"fmt"

	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
	"github.com/stretchr/testify/assert"
)

func (assignment *MimcCircuit) VerifyWitness(field *big.Int, publicWitness witness.Witness) bool {
	newWitness, err := assignment.ConstructWitness(field)
	if err != nil {
		return false
	}

	fmt.Println("public witness:", publicWitness)
	fmt.Println("constructed witness:", newWitness)

	return assert.Equal(&assert.CollectT{}, publicWitness, newWitness)
	//return publicWitness == newWitness
}

func VerifyWitness2(field *big.Int, publicWitness witness.Witness, publicHash any) bool {

	assignment := MimcCircuit{
		PreImage: 123,
		Hash:     publicHash,
	}

	witness, _ := frontend.NewWitness(&assignment, field)
	newPublicWitness, _ := witness.Public()
	fmt.Println("public witness:", publicWitness)
	fmt.Println("constructed witness:", newPublicWitness)

	return assert.Equal(&assert.CollectT{}, publicWitness, newPublicWitness)
}

// constructs new public witness using assignment's public inputs
func (assignment *MimcCircuit) ConstructWitness(field *big.Int) (witness.Witness, error) {
	newWitness, err := witness.New(field)
	if err != nil {
		return nil, err
	}

	witnessChan := make(chan any)
	go assignment.passPubInputs(&witnessChan)
	newWitness.Fill(1, 0, witnessChan)

	return newWitness, nil
}

// close the channel after passing the values to end the for loop over channel values
func (assignment *MimcCircuit) passPubInputs(witnessChan *chan any) {
	*witnessChan <- assignment.Hash
	fmt.Println("pulbic values sent via channel for witness construction...")
	close(*witnessChan)
}
