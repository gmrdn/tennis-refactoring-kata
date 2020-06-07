package tenniskata

type tennisGame3 struct {
	scorePlayer1 int
	scorePlayer2 int
	namePlayer1  string
	namePlayer2  string
}

// TennisGame3 creates a game with two player names
func TennisGame3(namePlayer1 string, namePlayer2 string) TennisGame {
	game := &tennisGame3{
		namePlayer1: namePlayer1,
		namePlayer2: namePlayer2}

	return game
}

func (game *tennisGame3) getTextualScore(points int) string {
	textualScores := []string{"Love", "Fifteen", "Thirty", "Forty"}
	return textualScores[points]
}

func (game *tennisGame3) isTie() bool {
	return game.scorePlayer1 == game.scorePlayer2
}

func (game *tennisGame3) isEarlyGame() bool {
	return game.scorePlayer1 < 4 && game.scorePlayer2 < 4 && !(game.scorePlayer1+game.scorePlayer2 == 6)
}

func (game *tennisGame3) whoIsLeading() string {
	if game.scorePlayer1 > game.scorePlayer2 {
		return game.namePlayer1
	}
	return game.namePlayer2

}

func (game *tennisGame3) pointsDifference() int {
	return (game.scorePlayer1 - game.scorePlayer2) * (game.scorePlayer1 - game.scorePlayer2)
}

func (game *tennisGame3) GetScore() string {
	if game.isEarlyGame() {
		gameScore := game.getTextualScore(game.scorePlayer1)
		if game.isTie() {
			return gameScore + "-All"
		}
		return gameScore + "-" + game.getTextualScore(game.scorePlayer2)
	}

	if game.isTie() {
		return "Deuce"
	}

	if game.pointsDifference() == 1 {
		return "Advantage " + game.whoIsLeading()
	}
	return "Win for " + game.whoIsLeading()

}

func (game *tennisGame3) WonPoint(playerName string) {
	if playerName == game.namePlayer1 {
		game.scorePlayer1++
	} else {
		game.scorePlayer2++
	}
}
