package comparisions

import (
	"errors"

	"github.com/consensys/gnark/frontend"
)

type ComparisionCircuit struct {
	Add  frontend.Variable
	Add2 frontend.Variable
}

func (circuit *ComparisionCircuit) Define(api frontend.API) error {

	// var x frontend.Variable
	// x = 0

	// for i := 0; i < 5; i++ {
	// 	toAdd := api.Select(circuit.Add, i, 0)
	// 	x = api.Add(x, toAdd)
	// }

	// api.AssertIsLessOrEqual(10, x)

	// var x frontend.Variable
	// // var y frontend.Variable

	// x = 2
	// y = 2

	// fmt.Println(x, y)
	//api.AssertIsEqual(x, circuit.Add)

	if circuit.Add2 != circuit.Add {
		// fmt.Println("x and y are", x, circuit.Add)
		return errors.New("not equal")
	}

	return nil
}
