package game

import (
	"math/rand"
)

// PlayWithOnePlayer returns a game result based on user's `choice` and a random system choice
func PlayWithOnePlayer(userChoice string) OnePlayerGameResult {
	result := OnePlayerGameResult{UserChoice: userChoice}

	systemChoice := AllowedChoices[rand.Intn(len(AllowedChoices))]

	result.SystemChoice = systemChoice

	if message, ok := choiceResultMap[ChoicePair{PlayerOneChoice: userChoice, PlayerTwoChoice: systemChoice}]; ok {
		result.DidUserWin = "WIN"
		result.Message = message
	} else if message, ok := choiceResultMap[ChoicePair{PlayerOneChoice: systemChoice, PlayerTwoChoice: userChoice}]; ok {
		result.DidUserWin = "LOSE"
		result.Message = message
	} else if userChoice == systemChoice {
		result.DidUserWin = "DRAW"
		result.Message = "Game drew. Try again."
	}

	return result
}

// PlayWithTwoPlayers returns a game result based on the two players' choices
func PlayWithTwoPlayers(p1Choice, p2Choice string) TwoPlayerGameResult {
	result := TwoPlayerGameResult{P1Choice: p1Choice, P2Choice: p2Choice}

	if message, ok := choiceResultMap[ChoicePair{PlayerOneChoice: p1Choice, PlayerTwoChoice: p2Choice}]; ok {
		result.Winner = 1
		result.Message = message
	} else if message, ok := choiceResultMap[ChoicePair{PlayerOneChoice: p2Choice, PlayerTwoChoice: p1Choice}]; ok {
		result.Winner = 2
		result.Message = message
	} else if p1Choice == p2Choice {
		result.Winner = 0
		result.Message = "Game drew. Try again."
	}

	return result
}
