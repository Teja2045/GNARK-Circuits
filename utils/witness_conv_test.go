package utils

import (
	"fmt"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

type IntCircuit struct {
	A frontend.Variable `gnark:",public"`
	B frontend.Variable
}

func (circuit *IntCircuit) Define(api frontend.API) error {
	api.Println("in circuit", circuit.A)
	api.Println("in circuit", circuit.B)
	api.AssertIsEqual(circuit.A, circuit.B)
	return nil
}

func TestInt(t *testing.T) {

	str := "21888242871839275222246405745257275088548364400416034343698204186575808495618" // Example string representation of a number
	str2 := "1"

	// // Convert string to a big.Int
	// a := new(big.Int)
	// _, success := a.SetString(str, 10) // Assuming base 10

	// if !success {
	// 	fmt.Println("Failed to convert string to bigint")
	// 	return
	// }

	assignment := IntCircuit{
		A: str,
		B: str2,
	}

	scalarField1 := ecc.BN254.ScalarField()
	scalarField2 := ecc.BW6_633.ScalarField()
	scalarField3 := ecc.BLS12_377.ScalarField()

	fmt.Println("fields")
	fmt.Println(scalarField1)
	fmt.Println(scalarField2)
	fmt.Println(scalarField3)

	witn, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())

	WitnessToInt(witn)
	fmt.Println(witn)
	fmt.Println()
	assrt := test.NewAssert(t)
	assrt.CheckCircuit(&IntCircuit{})
	assrt.ProverSucceeded(&IntCircuit{}, &assignment)
	// panic("done")

}
