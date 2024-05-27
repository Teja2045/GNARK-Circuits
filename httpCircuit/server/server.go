package main

import (
	"encoding/binary"
	"fmt"
	"net/http"
)

// an endpoint which returns different data when each time called will not work for ZK circuits !!!
func scoreHandler(w http.ResponseWriter, r *http.Request) {
	// rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	// randomNumber := uint64(rand.Intn(101) + 100)
	randomNumber := uint64(1000)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, randomNumber)
	w.Write(bytes)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to random score generator")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/score", scoreHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
