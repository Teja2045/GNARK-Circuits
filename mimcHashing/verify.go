package mimcHashing

import (
	"fmt"
	"zkCircuits/utils"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func Verify() {
	var circuit MimcCircuit
	ccs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	pk, vk, _ := groth16.Setup(ccs)

	dummyData := utils.GetDummyHashedData()
	assignment := MimcCircuit{
		PreImage: dummyData.Data,
		Hash:     utils.HexStringToByteArray(dummyData.HashString),
	}

	fmt.Println("Hash is: ", assignment.Hash)

	witness, _ := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())

	publicWitness, _ := witness.Public()

	//witness2, _ := frontend.NewWitness()

	big := ecc.BN254.ScalarField()
	fmt.Println("verify public witness: ", assignment.VerifyWitness(big, publicWitness))

	fmt.Println("hash string is ", dummyData.HashString)
	fmt.Println("hash by array is ", utils.HexStringToByteArray(dummyData.HashString))
	fmt.Println("------------ Public Witness is ", witness, "--------------")

	proof, _ := groth16.Prove(ccs, pk, witness)

	groth16.Verify(proof, vk, publicWitness)
}
