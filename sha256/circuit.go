package sha256

import (
	"fmt"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/uints"
)

type Sha256Circuit struct {
	In       []uints.U8
	Expected []uints.U8
}

func (circuit *Sha256Circuit) Define(api frontend.API) error {
	hFunc, err := New(api)
	if err != nil {
		return err
	}

	uapi, err := uints.New[uints.U32](api)
	if err != nil {
		return err
	}
	hFunc.Write(circuit.In)
	hash := hFunc.Sum()
	if len(hash) != 32 {
		return fmt.Errorf("not 32 bytes")
	}

	for i := range circuit.Expected {
		uapi.ByteAssertEq(circuit.Expected[i], hash[i])
	}
	return nil
}
