package signature

import (
	"testing"

	"github.com/Teja2045/GNARK-Circuits/utils"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/signature/eddsa"
	"github.com/consensys/gnark/test"
)

func TestSignatureCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var circuit SignatureCircuit

	mimc := hash.MIMC_BN254
	hFunc := mimc.New()

	// bigger than this is problem
	data := make([]byte, 32)
	for i := byte(0); i < 32; i++ {
		data[i] = i
	}

	privKey, pubKey := utils.GenerateKeys(0)
	signature := utils.Sign(data, privKey, hFunc)
	signBytes, err := privKey.Sign(data, hFunc)

	assert.NoError(err)

	hFunc.Reset()
	good, err := pubKey.Verify(signBytes, data, hFunc)

	assert.Equal(good, true)
	assert.NoError(err)

	var pbKey eddsa.PublicKey
	pbKey.A.X = pubKey.A.X
	pbKey.A.Y = pubKey.A.Y

	witness := &SignatureCircuit{
		PubKey: pbKey,
		Data:   data,
	}

	witness.Signature.R.X = signature.R.X
	witness.Signature.R.Y = signature.R.Y
	witness.Signature.S = signature.S[:]

	assert.ProverSucceeded(
		&SignatureCircuit{},
		witness,
		test.WithCurves(ecc.BN254),
		test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()),
	)

	assert.CheckCircuit(&circuit, test.WithValidAssignment(witness), test.WithCurves(ecc.BN254))

}
