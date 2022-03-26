/*
Package cli implements a command line interface to play a word game.

*/
package main

import (
	"fmt"
	"sort"
	"strings"

	gtw "github.com/gmofishsauce/gtw/lib"
)

const defaultCorpusName = "webster-2-all-five-letter.corpus"

const help = `
--------
After each guess, a signature will be displayed. In the signature,
the character '-' means the letter is not in the word. Lower case
letters are not in the right place, while upper case letters are
correctly placed.  Example:

guess> tears
       --ers (0 letters in the correct place)
guess> cloud
       -l-u- (0 letters in the correct place)
guess> aural
       *URAL (4 letters in the correct place)
guess> rural

Success!
--------
`

func main() {
	corpus, _ := gtw.LoadFile("cmd/cli/wordle.corpus")
	sort.Strings(corpus)

	benchmarkAnswers := []string{
		"epoxy",
		// "depot",
		// "chest",
		// "purge",
		// "slosh",
		"their",
		// "renew",
		// "allow",
		// "saute",
		// "movie",
		// "cater",
		// "tease",
		// "smelt",
		// "focus",
		// "today",
		"watch"}
	// "lapse",
	// "month",
	// "sweet",
	// "hoard",
	// "cloth",
	// "brine",
	// "ahead",
	// "mourn",
	// "nasty",
	// "rupee"}
	methods := []string{"dumb", "partition_size", "max_partition_size", "partition_size_deviation"}
	fmt.Printf("answer, %s\n", strings.Join(methods, ", "))
	for _, answer := range benchmarkAnswers {
		counts := make([]string, len(methods))
		for i, method := range methods {
			roundCount := 0
			agent := gtw.Agent(corpus)
			// fmt.Printf("Remaining word count is %d\n", agent.RemainingWordCount())
			for agent.RemainingWordCount() > 0 {
				guess := agent.GuessUsingMethod(method)
				signature, _ := gtw.ScoreAgainstGoal(guess, answer)
				nextAgent := agent.Inform(guess, signature)
				roundCount = roundCount + 1
				fmt.Printf("Answer %s method %s. Round %d: %s reduces from %d to %d.\n", answer, method, roundCount, guess, agent.RemainingWordCount(), nextAgent.RemainingWordCount())
				agent = nextAgent

				if guess == answer {
					break
				}
			}

			// fmt.Printf("Answer %s method %s -> rounds %d\n", answer, method, roundCount)
			counts[i] = fmt.Sprint(roundCount)
		}
		fmt.Printf("%s, %s\n", answer, strings.Join(counts, ", "))
	}
}
