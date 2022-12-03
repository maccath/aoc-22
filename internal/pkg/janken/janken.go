package janken

const Win string = "win"
const Lose string = "lose"
const Draw string = "draw"

const Rock string = "rock"
const Paper string = "paper"
const Scissors string = "scissors"

func Play(you string, opponent string) string {
	if you == opponent {
		return Draw
	}

	if you == Rock && opponent == Paper {
		return Lose
	}

	if you == Paper && opponent == Scissors {
		return Lose
	}

	if you == Scissors && opponent == Rock {
		return Lose
	}

	return Win
}

func Fix(p1 string, outcome string) string {
	if outcome == Win {
		if p1 == Rock {
			return Paper
		}
		if p1 == Paper {
			return Scissors
		}
		if p1 == Scissors {
			return Rock
		}
	}

	if outcome == Lose {
		if p1 == Rock {
			return Scissors
		}
		if p1 == Paper {
			return Rock
		}
		if p1 == Scissors {
			return Paper
		}
	}

	// Draw
	return p1
}

func Score(move string, outcome string) int {
	scores := make(map[string]int)
	scores[Rock] = 1
	scores[Paper] = 2
	scores[Scissors] = 3

	scores[Win] = 6
	scores[Lose] = 0
	scores[Draw] = 3

	return scores[move] + scores[outcome]
}
