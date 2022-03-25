package agent

import (
	"fmt"
	"strings"
	
	"github.com/gmofishsauce/gtw/lib"
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
		keep := true
		for i, letter := range guess {
			if signature[i] == gtw.LETTER_CORRECT {
				if rune(word[i]) != letter {
					keep = false
					break
				}
			} else if signature[i] == gtw.LETTER_IN_WORD {
				if !strings.Contains(word, string(letter)) {
					keep = false
					break
				}
			} else if signature[i] == gtw.LETTER_WRONG {
				if rune(word[i]) == letter {
					keep = false
					break
				}
			} else {
				panic(signature[i])
			}
		}

		if keep {
			matching = append(matching, word)
		}
	}

	fmt.Printf("Guess %s filtered corpus from %d words down to %d words\n", guess, len(agent.remaining), len(matching))

	return Agent(matching)
}
