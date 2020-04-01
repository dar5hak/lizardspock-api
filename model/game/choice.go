package game

const (
	// Rock crushes scissors and lizard
	Rock = "rock"

	// Paper covers rock and disproves Spock
	Paper = "paper"

	// Scissors cuts paper and decapitates lizard
	Scissors = "scissors"

	// Lizard eats paper and poisons Spock
	Lizard = "lizard"

	// Spock smashes scissors and vaporizes rock
	Spock = "spock"
)

// AllowedChoices are the only possible choices a user or system may make
var AllowedChoices = []string{Rock, Paper, Scissors, Lizard, Spock}

// ChoicePair is a tuple containing the choice made by both sides
type ChoicePair struct {
	PlayerOneChoice, PlayerTwoChoice string
}

var choiceResultMap map[ChoicePair]string = map[ChoicePair]string{
	ChoicePair{PlayerOneChoice: Scissors, PlayerTwoChoice: Paper}:  "Scissors cuts paper.",
	ChoicePair{PlayerOneChoice: Paper, PlayerTwoChoice: Rock}:      "Paper covers rock.",
	ChoicePair{PlayerOneChoice: Rock, PlayerTwoChoice: Lizard}:     "Rock crushes lizard.",
	ChoicePair{PlayerOneChoice: Lizard, PlayerTwoChoice: Spock}:    "Lizard poisons Spock.",
	ChoicePair{PlayerOneChoice: Spock, PlayerTwoChoice: Scissors}:  "Spock smashes scissors.",
	ChoicePair{PlayerOneChoice: Scissors, PlayerTwoChoice: Lizard}: "Scissors decapitates lizard.",
	ChoicePair{PlayerOneChoice: Lizard, PlayerTwoChoice: Paper}:    "Lizard eats paper.",
	ChoicePair{PlayerOneChoice: Paper, PlayerTwoChoice: Spock}:     "Paper disproves Spock.",
	ChoicePair{PlayerOneChoice: Spock, PlayerTwoChoice: Rock}:      "Spock vaporizes rock.",
	ChoicePair{PlayerOneChoice: Rock, PlayerTwoChoice: Scissors}:   "And as it always has, rock crushes scissors.",
}

