package main

import (
	"fmt"

	"github.com/Teja2045/GNARK-Circuits/cubic"
	"github.com/Teja2045/GNARK-Circuits/mimcHashing"
)

func main() {
	cubic.Verify()

	fmt.Println()

	mimcHashing.Verify()
}
