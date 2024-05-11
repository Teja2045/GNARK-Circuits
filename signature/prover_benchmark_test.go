package signature

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func BenchmarkProver(b *testing.B) {
	var circuit SignatureCircuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		b.Fatal(err)
	}

	pk, _, err := groth16.Setup(ccs)
	if err != nil {
		b.Fatal(err)
	}
	data, pbKey, sign := GetRandomSigntureData()

	b.ResetTimer() // reset timer before starting benchmark

	for i := 0; i < b.N; i++ {
		// Measure the time taken for Prover function
		_, err := Prover(pk, pbKey, data, sign)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkProver-12    	     576	   1817306 ns/op	  241759 B/op	    1195 allocs/op
// PASS
// ok  	github.com/Teja2045/GNARK-Circuits/signature	1.280s

func BenchmarkVerfier(b *testing.B) {
	var circuit SignatureCircuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		b.Fatal(err)
	}

	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		b.Fatal(err)
	}
	data, pbKey, sign := GetRandomSigntureData()

	proof, err := Prover(pk, pbKey, data, sign)
	if err != nil {
		b.Fatal(err)
	}

	_, pbKey2, sign2 := GetRandomSigntureData2()

	b.ResetTimer() // reset timer before starting benchmark

	for i := 0; i < b.N; i++ {
		// Measure the time taken for Verifier function
		Verifier(proof, vk, data, pbKey2, sign2)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkVerfier-12    	     547	   2426138 ns/op	   38796 B/op	     277 allocs/op
// PASS
// ok  	github.com/Teja2045/GNARK-Circuits/signature	1.573s
