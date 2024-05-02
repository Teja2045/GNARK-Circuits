package main

import (
	"fmt"
	"zkCircuits/cubic"
	"zkCircuits/mimcHashing"
)

func main() {
	cubic.Verify()

	fmt.Println()

	mimcHashing.Verify()
}
