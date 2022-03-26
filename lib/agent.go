package gtw

import (
	"fmt"
)

type GtwAgent struct {
	remaining []string
}

func Agent(corpus []string) *GtwAgent {
	return &GtwAgent{corpus}
}

func (agent *GtwAgent) Inform(guess string, signature string) *GtwAgent {
	matching := make([]string, 0, len(agent.remaining))

	for _, word := range agent.remaining {
		wordLetterCounts := make(map[rune]int)
		for _, letter := range word {
			count, exists := wordLetterCounts[rune(letter)]

			if !exists {
				count = 0
			}

			wordLetterCounts[rune(letter)] = count + 1
		}

		keep := true
		for i, letter := range guess {
			count, _ := wordLetterCounts[rune(letter)]
			if signature[i] == LETTER_CORRECT {
				wordLetterCounts[rune(letter)] = count - 1
				if rune(word[i]) != letter {
					keep = false
					break
				}
			} else if signature[i] == LETTER_WRONG {
				if rune(word[i]) == letter {
					keep = false
					break
				}
			}
		}

		if keep {
			for i, letter := range guess {
				if signature[i] == LETTER_IN_WORD {
					count, exists := wordLetterCounts[rune(letter)]
					if exists && count > 0 {
						wordLetterCounts[rune(letter)] = count - 1
					} else {
						keep = false
						break
					}
				}
			}
		}

		if keep {
			matching = append(matching, word)
		}
	}

	fmt.Printf("Guess %s filtered corpus from %d words down to %d words\n", guess, len(agent.remaining), len(matching))

	return Agent(matching)
}
