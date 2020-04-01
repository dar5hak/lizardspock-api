package game

// OnePlayerGameResult is the result of a game between a user and the system
type OnePlayerGameResult struct {
	UserChoice   string `json:"userChoice,omitempty"`
	SystemChoice string `json:"systemChoice,omitempty"`
	Message      string `json:"message,omitempty"`
	DidUserWin   string `json:"didYouWin,omitempty"`
}

// OnePlayerGameError is returned when /play/1p fails
type OnePlayerGameError struct {
	UserChoice   string `json:"userChoice,omitempty"`
	ErrorMessage string `json:"error"`
}

func (e *OnePlayerGameError) Error() string {
	return e.ErrorMessage
}

