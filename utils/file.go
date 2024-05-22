package utils

import (
	"bytes"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
)

// CheckFileExists checks if a file exists and is not a directory.
func CheckFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ReadProvingKey reads a ProvingKey from a file.
func ReadProvingKey(filename string) (groth16.ProvingKey, error) {
	pk := groth16.NewProvingKey(ecc.BN254)
	data, err := os.ReadFile(filename)
	if err != nil {
		return pk, err
	}
	reader := bytes.NewReader(data)
	_, err = pk.ReadFrom(reader)
	return pk, err
}

// ReadVerifyingKey reads a VerifyingKey from a file.
func ReadVerifyingKey(filename string) (groth16.VerifyingKey, error) {
	vk := groth16.NewVerifyingKey(ecc.BN254)
	data, err := os.ReadFile(filename)
	if err != nil {
		return vk, err
	}
	reader := bytes.NewReader(data)
	_, err = vk.ReadFrom(reader)
	return vk, err
}

// ReadConstraintSystem reads a ConstraintSystem from a file.
func ReadConstraintSystem(filename string) (constraint.ConstraintSystem, error) {
	cs := groth16.NewCS(ecc.BN254)
	data, err := os.ReadFile(filename)
	if err != nil {
		return cs, err
	}
	reader := bytes.NewReader(data)
	_, err = cs.ReadFrom(reader)

	return cs, err
}

// WriteProvingKey writes a ProvingKey to a file.
func WriteProvingKey(filename string, pk groth16.ProvingKey) error {
	var buf bytes.Buffer
	_, err := pk.WriteTo(&buf)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

// WriteVerifyingKey writes a VerifyingKey to a file.
func WriteVerifyingKey(filename string, vk groth16.VerifyingKey) error {
	var buf bytes.Buffer
	_, err := vk.WriteTo(&buf)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}

// WriteConstraintSystem writes a ConstraintSystem to a file.
func WriteConstraintSystem(filename string, cs constraint.ConstraintSystem) error {
	var buf bytes.Buffer
	_, err := cs.WriteTo(&buf)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buf.Bytes(), 0644)
}
