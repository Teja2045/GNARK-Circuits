package ipl

import (
	"github.com/consensys/gnark/frontend"
)

type IplQualierCircuit struct {
	Team1RunRate    frontend.Variable `gnark:",public"`
	Team2RunRate    frontend.Variable
	Team1MatchScore frontend.Variable
	Team2MatchScore frontend.Variable
}

// a naive ipl circuit, just for testing circuits

func (circuit *IplQualierCircuit) Define(api frontend.API) error {

	matchScoreDiff := api.Sub(circuit.Team1MatchScore, circuit.Team2MatchScore)
	matchRunRateDiff := api.Div(matchScoreDiff, 10)
	team1FinalRunRate := api.Add(circuit.Team1RunRate, matchRunRateDiff)
	team2FinalRunRate := api.Sub(circuit.Team2RunRate, matchRunRateDiff)
	//fmt.Println("DDDDDDDDDDDDD", team1FinalRunRate, team2FinalRunRate)
	api.AssertIsLessOrEqual(team2FinalRunRate, team1FinalRunRate)

	return nil
}
