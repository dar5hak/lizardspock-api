package game

// TwoPlayerGameResult is the result of a game between two players
type TwoPlayerGameResult struct {
	P1Choice string `json:"p1Choice,omitempty"`
	P2Choice string `json:"p2Choice,omitempty"`
	Message  string `json:"message,omitempty"`
	Winner   uint8  `json:"winner,omitempty"`
}

// TwoPlayerGameError is returned when /play/2p fails
type TwoPlayerGameError struct {
	P1Choice     string `json:"p1Choice,omitempty"`
	P2Choice     string `json:"p2Choice,omitempty"`
	ErrorMessage string `json:"error"`
}

func (e *TwoPlayerGameError) Error() string {
	return e.ErrorMessage
}

