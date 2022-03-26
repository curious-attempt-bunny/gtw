package gtw

import "fmt"

type GtwAgent struct {
	remaining []string
}

func Agent(corpus []string) *GtwAgent {
	return &GtwAgent{corpus}
}

func (agent *GtwAgent) RemainingWordCount() int {
	return len(agent.remaining)
}

func (agent *GtwAgent) Inform(guess string, signature string) *GtwAgent {
	matching := make([]string, 0, len(agent.remaining))

	guessCorrectLetterCounts := make(map[rune]int)
	for i, letter := range guess {
		if signature[i] != LETTER_WRONG {
			count, exists := guessCorrectLetterCounts[rune(letter)]

			if !exists {
				count = 0
			}

			guessCorrectLetterCounts[rune(letter)] = count + 1
		}
	}

	for _, word := range agent.remaining {
		wordLetterCounts := make(map[rune]int)
		for _, letter := range word {
			count, exists := wordLetterCounts[rune(letter)]

			if !exists {
				count = 0
			}

			wordLetterCounts[letter] = count + 1
		}

		keep := true
		for i, letter := range guess {
			if signature[i] == LETTER_CORRECT {
				fmt.Printf("Word %s letter %d %v: correct\n", word, i, letter)
				if rune(word[i]) != rune(letter) {
					// if guess == "genre" {
					fmt.Printf("Removed %v due to letter %d. %v is correct and != %v\n", word, i, rune(letter), word[i])
					// }
					keep = false
					break
				}
			} else if signature[i] == LETTER_WRONG {
				if rune(word[i]) == rune(letter) {
					// if guess == "genre" {
					fmt.Printf("Removed %v due to letter %d. %v is wrong and == %v\n", word, i, letter, word[i])
					// }
					keep = false
					break
				}
			} else if signature[i] == LETTER_IN_WORD {
				guessCount, guessExists := guessCorrectLetterCounts[rune(letter)]
				wordCount, wordExists := wordLetterCounts[rune(letter)]

				if !guessExists {
					panic("Expected guessCorrectLetterCounts to contain letter!")
				}

				if !wordExists || wordCount < guessCount {
					fmt.Printf("Removed %v due to having insufficient floating letter %d. %v is floating %d times.\n", word, i, letter, wordCount)
					keep = false
					break
				} else {
					fmt.Printf("Kept %v due to having sufficient floating letter %d. %v is floating %d times.\n", word, i, letter, wordCount)
				}
			}
		}

		// if keep && len(wrongPlaceLetterCounts) > 0 {
		// 	if guess == "newer" {
		// 		fmt.Printf("Considering words in wrong places for %s with %v.\n", word, wrongPlaceLetterCounts)
		// 	}
		// 	for i, letter := range guess {
		// 		if signature[i] == LETTER_IN_WORD {
		// 			count, exists := wrongPlaceLetterCounts[rune(letter)]
		// 			if exists {
		// 				if count > 0 {
		// 					wordLetterCounts[rune(letter)] = count - 1
		// 					if guess == "newer" {
		// 						fmt.Printf("Still keeping %v due to letter %d. %v is in the wrong place and count is %d\n", word, i, letter, count)
		// 					}
		// 				} else {
		// 					if guess == "newer" {
		// 						fmt.Printf("Removing %v due to letter %d. %v is in the wrong place and count is %d\n", word, i, letter, count)
		// 					}
		// 					keep = false
		// 					break
		// 				}
		// 			}
		// 		}
		// 	}
		// }

		if guess == word {
			keep = false
		}

		if keep {
			matching = append(matching, word)
		}
	}

	// fmt.Printf("Guess %s filtered corpus from %d words down to %d words\n", guess, len(agent.remaining), len(matching))

	if len(matching) == len(agent.remaining) && len(matching) > 1 {
		fmt.Printf("guess %s and %v = %v\n", guess, matching, agent.remaining)
		panic("Impossible!")
	}
	return Agent(matching)
}
