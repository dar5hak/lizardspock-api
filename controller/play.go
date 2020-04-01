package controller

import (
	"github.com/gin-gonic/gin"
	"lizardspock/model/game"
	"lizardspock/util"
	"net/http"
	"strings"
)

// Play1p godoc
// @Summary Play against the system
// @Tags play
// @Accept  json
// @Produce  json
// @param choice query string true "your choice: one of rock, paper, scissors, lizard or spock"
// @Success 200 {object} game.OnePlayerGameResult
// @Failure 400 {object} game.OnePlayerGameError
// @Router /play/1p [get]
func Play1p(c *gin.Context) {
	choice, exists := c.GetQuery("choice")
	if !exists {
		c.JSON(http.StatusBadRequest, game.OnePlayerGameError{ErrorMessage: "no choice made"})
	}

	if !util.Contains(game.AllowedChoices, strings.ToLower(choice)) {
		c.JSON(http.StatusBadRequest, game.OnePlayerGameError{UserChoice: choice, ErrorMessage: "invalid choice"})
	}

	c.JSON(http.StatusOK, game.PlayWithOnePlayer(choice))
}

// Play2p godoc
// @Summary Play a two-player game
// @Tags play
// @Accept  json
// @Produce  json
// @param p1Choice query string true "player 1's choice: one of rock, paper, scissors, lizard or spock"
// @param p2Choice query string true "player 2's choice: one of rock, paper, scissors, lizard or spock"
// @Success 200 {object} game.TwoPlayerGameResult
// @Failure 400 {object} game.TwoPlayerGameError
// @Router /play/2p [get]
func Play2p(c *gin.Context) {
	p1Choice, exists := c.GetQuery("p1Choice")
	if !exists {
		c.JSON(http.StatusBadRequest, game.TwoPlayerGameError{ErrorMessage: "no choice made by player 1"})
	}

	p2Choice, exists := c.GetQuery("p2Choice")
	if !exists {
		c.JSON(http.StatusBadRequest, game.TwoPlayerGameError{ErrorMessage: "no choice made by player 2"})
	}

	if !util.Contains(game.AllowedChoices, strings.ToLower(p1Choice)) {
		c.JSON(http.StatusBadRequest, game.TwoPlayerGameError{P1Choice: p1Choice, P2Choice: p2Choice, ErrorMessage: "invalid player 1 choice"})
	} else if !util.Contains(game.AllowedChoices, strings.ToLower(p2Choice)) {
		c.JSON(http.StatusBadRequest, game.TwoPlayerGameError{P1Choice: p1Choice, P2Choice: p2Choice, ErrorMessage: "invalid player 2 choice"})
	}

	result := game.PlayWithTwoPlayers(p1Choice, p2Choice)

	c.JSON(http.StatusOK, result)
}
