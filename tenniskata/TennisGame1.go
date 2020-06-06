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

func (game *tennisGame1) getEqualScore(score int) string {
	textualScore := ""
	switch score {
	case 0:
		textualScore = "Love-All"
	case 1:
		textualScore = "Fifteen-All"
	case 2:
		textualScore = "Thirty-All"
	default:
		textualScore = "Deuce"
	}
	return textualScore
}


func (game *tennisGame1) GetScore() string {
	score := ""
	tempScore := 0
	if game.isEqualScore() {
		score = game.getEqualScore(game.mScore1)
	} else if game.mScore1 >= 4 || game.mScore2 >= 4 {
		minusResult := game.mScore1 - game.mScore2
		if minusResult == 1 {
			score = "Advantage player1"
		} else if minusResult == -1 {
			score = "Advantage player2"
		} else if minusResult >= 2 {
			score = "Win for player1"
		} else {
			score = "Win for player2"
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.mScore1
			} else {
				score += "-"
				tempScore = game.mScore2
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}