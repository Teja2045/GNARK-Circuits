package my_ed25519

import (
	"testing"

	"github.com/Teja2045/GNARK-Circuits/utils"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/signature/eddsa"
	"github.com/consensys/gnark/test"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
)

// not completed since 25519 curve is not implemented
func TestSignatureCircuit(t *testing.T) {
	assert := test.NewAssert(t)

	var circuit Ed25519Circuit

	mimc := hash.MIMC_BN254
	hFunc := mimc.New()

	data := []byte{1, 2, 3}

	secret := []byte("huge huge secret")
	sdkPrivkey := ed25519.GenPrivKeyFromSecret(secret)

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

	witness := &Ed25519Circuit{
		PubKey: pbKey,
		Data:   data,
	}

	witness.Signature.R.X = signature.R.X
	witness.Signature.R.Y = signature.R.Y
	witness.Signature.S = signature.S[:]

	assert.ProverSucceeded(
		&Ed25519Circuit{},
		witness,
		test.WithCurves(ecc.BN254),
		test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()),
	)

	assert.CheckCircuit(&circuit, test.WithValidAssignment(witness), test.WithCurves(ecc.BN254))

}
