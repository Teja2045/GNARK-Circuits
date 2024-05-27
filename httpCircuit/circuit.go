package httpCircuit

import (
	"github.com/Teja2045/GNARK-Circuits/httpCircuit/client"
	"github.com/consensys/gnark/frontend"
)

type HttpCircuit struct {
	Score frontend.Variable `gnark:",public"`
}

func (circuit *HttpCircuit) Define(api frontend.API) error {
	apiScore, err := client.GetScoreFromAPI()
	if err != nil {
		return err
	}

	// fmt.Println("printing api score", &apiScore)
	api.AssertIsLessOrEqual(apiScore, circuit.Score)
	api.AssertIsDifferent(circuit.Score, uint64(1000))
	return nil
}
