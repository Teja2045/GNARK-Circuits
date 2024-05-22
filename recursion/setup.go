package recursion

import (
	"fmt"
	"log/slog"

	"time"

	"github.com/Teja2045/GNARK-Circuits/utils"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/algebra/emulated/sw_bn254"
	stdgroth16 "github.com/consensys/gnark/std/recursion/groth16"
)

func InnerSetup() (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {
	start := time.Now()
	defer func(before time.Time) {
		timeTaken := time.Since(before).Seconds()
		slog.Info(fmt.Sprintln("time taken for inner setup: ", timeTaken, "seconds"))
	}(start)
	var circuit InnerCircuit
	return Setup(&circuit)
}

func PersistantInnerSetup() (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {
	start := time.Now()
	defer func(before time.Time) {
		timeTaken := time.Since(before).Seconds()
		slog.Info(fmt.Sprintf("time taken for outer setup: %.2f seconds", timeTaken))
	}(start)

	pkFile := "innerProvingKey.bin"
	vkFile := "innerVerifyingKey.bin"
	csFile := "innerConstraintSystem.bin"

	var pk groth16.ProvingKey
	var vk groth16.VerifyingKey
	var cs constraint.ConstraintSystem

	if utils.CheckFileExists(pkFile) && utils.CheckFileExists(vkFile) && utils.CheckFileExists(csFile) {
		var err error
		pk, err = utils.ReadProvingKey(pkFile)
		if err != nil {
			return pk, vk, cs, err
		}

		vk, err = utils.ReadVerifyingKey(vkFile)
		if err != nil {
			return pk, vk, cs, err
		}

		cs, err = utils.ReadConstraintSystem(csFile)
		if err != nil {
			return pk, vk, cs, err
		}

		slog.Info("succeess reading from files...")

		return pk, vk, cs, nil
	}

	var circuit InnerCircuit

	pk, vk, cs, err := Setup(&circuit)

	if err != nil {
		return pk, vk, cs, err
	}

	if err = utils.WriteProvingKey(pkFile, pk); err != nil {
		slog.Error("error in inner writing proving key to file..")
		return pk, vk, cs, nil
	}

	if err = utils.WriteVerifyingKey(vkFile, vk); err != nil {
		slog.Error("error in writing inner verifying key to file..")
		return pk, vk, cs, nil
	}

	if err = utils.WriteConstraintSystem(csFile, cs); err != nil {
		slog.Error("error in writing inner contraint system to file..")
		return pk, vk, cs, nil
	}

	return pk, vk, cs, nil
}

func OuterSetup(innerCcs constraint.ConstraintSystem) (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {
	start := time.Now()
	defer func(before time.Time) {
		timeTaken := time.Since(before).Seconds()
		slog.Info(fmt.Sprintln("time taken for outer setup: ", timeTaken, "seconds"))
	}(start)
	outerCircuit := &OuterCircuit[sw_bn254.ScalarField, sw_bn254.G1Affine, sw_bn254.G2Affine, sw_bn254.GTEl]{
		InnerWitness: stdgroth16.PlaceholderWitness[sw_bn254.ScalarField](innerCcs),
		VerifyingKey: stdgroth16.PlaceholderVerifyingKey[sw_bn254.G1Affine, sw_bn254.G2Affine, sw_bn254.GTEl](innerCcs),
	}
	pk, vk, cs, err := Setup(outerCircuit)

	return pk, vk, cs, err
}

func PersistantOuterSetup(innerCcs constraint.ConstraintSystem) (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {
	start := time.Now()
	defer func(before time.Time) {
		timeTaken := time.Since(before).Seconds()
		slog.Info(fmt.Sprintf("time taken for outer setup: %.2f seconds", timeTaken))
	}(start)

	pkFile := "outerProvingKey.bin"
	vkFile := "outerVerifyingKey.bin"
	csFile := "outerConstraintSystem.bin"

	var pk groth16.ProvingKey
	var vk groth16.VerifyingKey
	var cs constraint.ConstraintSystem

	if utils.CheckFileExists(pkFile) && utils.CheckFileExists(vkFile) && utils.CheckFileExists(csFile) {
		var err error
		pk, err = utils.ReadProvingKey(pkFile)
		if err != nil {
			return pk, vk, cs, err
		}

		vk, err = utils.ReadVerifyingKey(vkFile)
		if err != nil {
			return pk, vk, cs, err
		}

		cs, err = utils.ReadConstraintSystem(csFile)
		if err != nil {
			return pk, vk, cs, err
		}

		slog.Info("succeess reading from files...")

		return pk, vk, cs, nil
	}

	outerCircuit := &OuterCircuit[sw_bn254.ScalarField, sw_bn254.G1Affine, sw_bn254.G2Affine, sw_bn254.GTEl]{
		InnerWitness: stdgroth16.PlaceholderWitness[sw_bn254.ScalarField](innerCcs),
		VerifyingKey: stdgroth16.PlaceholderVerifyingKey[sw_bn254.G1Affine, sw_bn254.G2Affine, sw_bn254.GTEl](innerCcs),
	}

	pk, vk, cs, err := Setup(outerCircuit)
	if err != nil {
		return pk, vk, cs, err
	}

	if err = utils.WriteProvingKey(pkFile, pk); err != nil {
		slog.Error("error in writing outer proving key to file..")
		return pk, vk, cs, nil
	}

	if err = utils.WriteVerifyingKey(vkFile, vk); err != nil {
		slog.Error("error in writing outer verifying key to file..")
		return pk, vk, cs, nil
	}

	if err = utils.WriteConstraintSystem(csFile, cs); err != nil {
		slog.Error("error in writing outer contraint system to file..")
		return pk, vk, cs, nil
	}

	return pk, vk, cs, nil
}

func Setup(circuit frontend.Circuit) (groth16.ProvingKey, groth16.VerifyingKey, constraint.ConstraintSystem, error) {

	cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("proof setup failed: %w", err)
	}

	pk, vk, err := groth16.Setup(cs)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("proof setup failed: %w", err)
	}

	return pk, vk, cs, nil

}
