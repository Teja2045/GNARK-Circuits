package secp256k1_ecdsa

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/consensys/gnark/std/signature/ecdsa"
)

type Secp256k1Circuit[T, S emulated.FieldParams] struct {
	Signature ecdsa.Signature[S]
	Message   emulated.Element[S]
	Pubkey    ecdsa.PublicKey[T, S]
}

func (circuit *Secp256k1Circuit[T, S]) Define(api frontend.API) error {
	circuit.Pubkey.Verify(api, sw_emulated.GetCurveParams[T](), &circuit.Message, &circuit.Signature)
	return nil
}
