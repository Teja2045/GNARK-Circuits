package utils

import (
	"encoding/hex"
)

func ByteArrayToHexString(bytedata []byte) string {
	return hex.EncodeToString(bytedata)
}

func HexStringToByteArray(data string) []byte {
	byteArray, err := hex.DecodeString(data)
	if err != nil {
		panic("couldb't convert string to byte array" + err.Error())
	}
	return byteArray
}
