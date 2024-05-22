package ipl

import (
	"fmt"
	"log"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func ProveAndVerify() {
	var circuit IplQualierCircuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		log.Fatal(err)
	}

	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		log.Fatal(err)
	}

	assignment := IplQualierCircuit{
		Team1RunRate:    25,
		Team2RunRate:    20,
		Team1MatchScore: 190,
		Team2MatchScore: 120,
	}

	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}

	publicWitness, err := witness.Public()
	if err != nil {
		panic(err)
	}

	fmt.Println("witness", &publicWitness)

	proof, _ := groth16.Prove(ccs, pk, witness)

	//verifier
	if err := groth16.Verify(proof, vk, publicWitness); err != nil {
		log.Println("proof is not verified....", err)
	} else {
		log.Println("proof is verified....")
	}
}
