package tenniskata

type tennisGame1 struct {
	mScore1    int
	mScore2    int
	player1Name string
	player2Name string
}

// TennisGame1 Starts a tennis game with two players names
func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func getTextualScore(score int) string {
	textualScores := []string{"Love", "Fifteen", "Thirty", "Forty"}
	return textualScores[score]
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.mScore1++
	} else {
		game.mScore2++
	}
}

func (game *tennisGame1) isEqualScore() bool {
	return game.mScore1 == game.mScore2
}
func (game *tennisGame1) isEndGame() bool {
	return game.mScore1 >= 4 || game.mScore2 >= 4
}

func (game *tennisGame1) getEqualScore(score int) string {
	textualScore := ""
	if score < 3 {
		textualScore = getTextualScore(score) + "-All"
	} else {
		textualScore = "Deuce"
	}
	return textualScore
}

func (game *tennisGame1) getEndGameScore() string {
	scoreDifference := game.mScore1 - game.mScore2
	score := ""
	if scoreDifference == 1 {
		score = "Advantage player1"
	} else if scoreDifference == -1 {
		score = "Advantage player2"
	} else if scoreDifference >= 2 {
		score = "Win for player1"
	} else {
		score = "Win for player2"
	}
	return score
}

func (game *tennisGame1) getEarlyGameScore() string {
	score := getTextualScore(game.mScore1) + "-" + getTextualScore(game.mScore2)
	return score
}

func (game *tennisGame1) GetScore() string {
	score := ""
	if game.isEqualScore() {
		score = game.getEqualScore(game.mScore1)
	} else if game.isEndGame() {
		score = game.getEndGameScore()
	} else {
		score = game.getEarlyGameScore()
	}
	return score
}