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
	gameScore := ""
	if score < 3 {
		gameScore = getTextualScore(score) + "-All"
	} else {
		gameScore = "Deuce"
	}
	return gameScore
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (game *tennisGame1) getLeadingPlayer() string {
	if game.mScore1 - game.mScore2 > 0 {
		return game.player1Name
	} 
	return game.player2Name
}

func (game *tennisGame1) getAbsoluteScoreDifference() int {
	scoreDifference := game.mScore1 - game.mScore2
	if scoreDifference < 0 {
		return -scoreDifference
	}
	return scoreDifference
} 

func (game *tennisGame1) getEndGameScore() string {
	score := ""
	if game.getAbsoluteScoreDifference() == 1 {
		score = "Advantage " + game.getLeadingPlayer()
	} else {
		score = "Win for " + game.getLeadingPlayer()
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