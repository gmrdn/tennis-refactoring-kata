package tenniskata

type tennisGame2 struct {
	P1point int
	P2point int

	P1res       string
	P2res       string
	player1Name string
	player2Name string
}

func TennisGame2(player1Name string, player2Name string) TennisGame {
	game := &tennisGame2{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame2) isEquality() bool {
	return game.P1point == game.P2point
}

func (game *tennisGame2) isEarlyGameP1() bool {
	return game.P1point < 4
}

func (game *tennisGame2) isEarlyGameP2() bool {
	return game.P2point < 4
}


func (game *tennisGame2) isLateGameP1() bool  {
	return game.P1point >= 3
}

func (game *tennisGame2) isLateGameP2() bool  {
	return game.P2point >= 3
}


func (game *tennisGame2) isAdvantageP1() bool  {
	return game.P1point > game.P2point
}

func (game *tennisGame2) isAdvantageP2() bool  {
	return game.P1point < game.P2point
}


func (game *tennisGame2) getTextualScore(point int) string {
	textualScores := []string{"Love", "Fifteen", "Thirty", "Forty"}
	return textualScores[point]
}

func (game *tennisGame2) GetScore() string {
	score := ""
	if game.isEquality() && game.isEarlyGameP1() {
		score = game.getTextualScore(game.P1point)
		score += "-All"
	}

	if game.isEquality() && game.isLateGameP1() {
		score = "Deuce"
	}

	if game.isAdvantageP1() && game.isEarlyGameP1() {
		game.P1res = game.getTextualScore(game.P1point)
		game.P2res = game.getTextualScore(game.P2point)
		score = game.P1res + "-" + game.P2res
	}

	if game.isAdvantageP2() && game.isEarlyGameP2() {
		game.P1res = game.getTextualScore(game.P1point)
		game.P2res = game.getTextualScore(game.P2point)
		score = game.P1res + "-" + game.P2res
	}

	if game.isAdvantageP1() && game.isLateGameP2() {
		score = "Advantage " + game.player1Name
	}

	if game.isAdvantageP2() && game.isLateGameP1() {
		score = "Advantage " + game.player2Name
	}

	if game.P1point >= 4 && game.P2point >= 0 && (game.P1point-game.P2point) >= 2 {
		score = "Win for " + game.player1Name
	}
	if game.P2point >= 4 && game.P1point >= 0 && (game.P2point-game.P1point) >= 2 {
		score = "Win for " + game.player2Name
	}
	return score
}

func (game *tennisGame2) WonPoint(player string) {
	if player == game.player1Name {
		game.P1point++
	} else {
		game.P2point++
	}
}