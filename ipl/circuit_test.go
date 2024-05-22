package ipl

import (
	"testing"

	"github.com/consensys/gnark/test"
)

func TestIPLCircuit(t *testing.T) {

	assert := test.NewAssert(t)
	team1Runrate := 20
	team2Runrate := 25
	team1Score := 190
	team2Score := 170

	var circuit IplQualierCircuit
	//assert.CheckCircuit(&circuit)

	assignment := IplQualierCircuit{
		Team1RunRate:    team1Runrate,
		Team2RunRate:    team2Runrate,
		Team1MatchScore: team1Score,
		Team2MatchScore: team2Score,
	}

	assert.ProverSucceeded(&circuit, &assignment)

}
