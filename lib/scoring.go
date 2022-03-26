package gtw

const LETTER_CORRECT = '+' // This letter is correct and in position
const LETTER_IN_WORD = '*' // This letter is in the word, but out of position
const LETTER_WRONG = '#'   // This letter is not in the word at any position
const LETTER_INVALID = 0   // This can't ever occur in a guess or a goal

// Score returns two values indicating the goodness of a guess.
// The first return value is a string describing the match result
// for the guess. In this string, '+' means the letter is in the
// correct position in the goal word, '*' indicates the letter is
// in the goal word but not in the correct position, and '#' means
// the letter is not in the word. The integer value is the number
// of '+' characters in the match result string. Note: the function
// Humanize(signature, guess) can be used to produce a result string
// is easier for humans to read from the result of this method.

func ScoreAgainstGoal(guess string, goal string) (string, int) {
	var aGuess, aGoal, signature [5]rune

	for i, _ := range guess {
		aGuess[i] = rune(guess[i])
		aGoal[i] = rune(goal[i])
		signature[i] = LETTER_WRONG
	}

	// First find all the correct matches. Once found, they
	// play no further role in matching either in the goal
	// or in the guess. Then make a second pass over the guess
	// and score any letter that still exists in the goal as
	// an out-of-place letter.

	unsolvedLetterCounts := make(map[rune]int)

	nCorrect := 0
	for i, g := range aGuess {
		if g == aGoal[i] {
			aGuess[i] = 0
			aGoal[i] = 0
			signature[i] = LETTER_CORRECT
			nCorrect++
		} else {
			count, exists := unsolvedLetterCounts[aGoal[i]]

			if !exists {
				count = 0
			}

			unsolvedLetterCounts[aGoal[i]] = count + 1
		}
	}

	for i, g := range aGuess {
		if aGuess[i] != 0 {
			count, exists := unsolvedLetterCounts[g]
			if exists && count > 0 {
				unsolvedLetterCounts[g] = count - 1
				signature[i] = LETTER_IN_WORD
			}
		}
	}

	return string(signature[:]), nCorrect
}
