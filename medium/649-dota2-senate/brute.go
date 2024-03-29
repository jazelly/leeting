package Dota2Senate

import "strings"

func predictPartyVictory(senate string) string {
	// Count of Each Type of Senator to check for Winner
	rCount := strings.Count(senate, "R")
	dCount := len(senate) - rCount

	// To mark Banned Senators
	banned := make([]bool, len(senate))

	// Ban the candidate "toBan", immediate next to "startAt"
	ban := func(toBan byte, startAt int) {
		// Find the next eligible senator of "toBan" type
		// On found, mark him as banned
		for {
			if senate[startAt] == toBan && !banned[startAt] {
				banned[startAt] = true
				break
			}
			startAt = (startAt + 1) % len(senate)
		}
	}

	// Turn of Senator at this Index
	turn := 0

	// While both parties have at least one senator
	for rCount > 0 && dCount > 0 {

		if !banned[turn] {
			if senate[turn] == 'R' {
				ban('D', (turn+1)%len(senate))
				dCount--
			} else {
				ban('R', (turn+1)%len(senate))
				rCount--
			}
		}

		turn = (turn + 1) % len(senate)
	}

	// Return Winner depending on the count of each party
	if dCount == 0 {
		return "Radiant"
	}
	return "Dire"
}
