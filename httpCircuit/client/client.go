package client

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
)

func GetScoreFromAPI() (uint64, error) {

	resp, err := http.Get("http://localhost:8080/score")
	if err != nil {
		fmt.Println("Failed to make request:", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return 0, err
	}

	if len(body) != 8 {
		return 0, fmt.Errorf("invalid response length: expected 8 bytes, got %d bytes", len(body))
	}

	value := binary.BigEndian.Uint64(body)
	fmt.Println(value)

	//var score fr.Element

	fmt.Println("Response from server:", value)
	return value, nil
}
