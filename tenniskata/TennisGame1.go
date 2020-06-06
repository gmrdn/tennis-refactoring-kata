package tenniskata

const scoreZero = "Love"
const scoreOne = "Fifteen"
const scoreTwo = "Thirty"
const scoreThree = "Forty"

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
	switch score {
	case 0:
		textualScore = scoreZero + "-All"
	case 1:
		textualScore = scoreOne + "-All"
	case 2:
		textualScore = scoreTwo + "-All"
	default:
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
	score := ""
	tempScore := 0
	for i := 1; i < 3; i++ {
		if i == 1 {
			tempScore = game.mScore1
		} else {
			score += "-"
			tempScore = game.mScore2
		}
		switch tempScore {
		case 0:
			score += scoreZero
		case 1:
			score += scoreOne
		case 2:
			score += scoreTwo
		case 3:
			score += scoreThree
		}
	}
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