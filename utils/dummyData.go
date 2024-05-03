package utils

import "math/big"

type HashedData struct {
	Data       big.Int
	HashString string
}

// The BN254 Mimc HashString of bigInt(10) is: 2f81d93f7e87768fcb6ae2ad6c782cf95a0d9dfee24bed956a85f7726f57b839
func GetDummyHashedData() HashedData {

	// or we can use bytes durectly also
	// but can't use string as frontend.Variable in circuit define method

	return HashedData{*big.NewInt(10), "2f81d93f7e87768fcb6ae2ad6c782cf95a0d9dfee24bed956a85f7726f57b839"}
}
