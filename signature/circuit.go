package signature

import (
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"

	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/consensys/gnark/std/signature/eddsa"
)

type SignatureCircuit struct {
	PubKey    eddsa.PublicKey
	Signature eddsa.Signature
	Data      frontend.Variable `gnark:",public"`
}

func (circuit *SignatureCircuit) Define(api frontend.API) error {
	hFunc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	curve, err := twistededwards.NewEdCurve(api, tedwards.BN254)
	if err != nil {
		return err
	}

	hFunc.Reset()
	err = eddsa.Verify(curve, circuit.Signature, circuit.Data, circuit.PubKey, &hFunc)
	if err != nil {
		return err
	}

	return nil
}
