/*
Package gtw implements a word game.

Artifacts in this package are suitable for use when
implementing a user interface for the word game or
when creating bots to play the word game.

*/
package gtw

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// LoadFile loads a "corpus file" having one word per newline-separated
// line and returns the file as an array of strings, one per line.
func LoadFile(filepath string) ([]string, error) {
	words, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(words)

	wordlist := make([]string, 0, 100)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return wordlist, nil
}

// GtwEngine is a "game engine" for Guess the Word
type GtwEngine struct {
	corpus []string
	rng    *rand.Rand
	goal   string
}

// New creates a new GtW evaluation engine given a corpus of words.
// The corpus may be constructed by LoadFile.
func New(corpus []string) *GtwEngine {
	if len(corpus) == 0 {
		panic("0-length corpus ... ouch, don't do that")
	}
	result := &GtwEngine{corpus, nil, ""}
	result.SetSeed(-1) // random
	result.NewGame()
	return result
}

// Get the Corpus
func (e *GtwEngine) Corpus() []string {
	return e.corpus
}

// Set the seed for the RNG
func (e *GtwEngine) SetSeed(seed int64) {
	if seed < 0 {
		seed = time.Now().UnixNano()
	}
	e.rng = rand.New(rand.NewSource(seed))
}

// NewGame reinitializes the goal word of the engine to a uniformly-
// selected random word from the engine's corpus.
func (e *GtwEngine) NewGame() {
	e.goal = e.corpus[e.rng.Int31n(int32(len(e.corpus)))]
}

// NewFixedGame reinitializes the goal word to the argument. The
// argument value must be in the corpus.
func (e *GtwEngine) NewFixedGame(aWord string) error {
	// If speed is of the essence,
	// this sanity code can be removed.
	found := false
	for _, v := range e.corpus {
		if v == aWord {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("candidate goal word not in corpus: %s", aWord)
	}
	e.goal = aWord
	return nil
}

// Cheat returns the the engine's current goal word.
func (e *GtwEngine) Cheat() string {
	return e.goal
}

func (e *GtwEngine) Score(guess string) (string, int) {
	return ScoreAgainstGoal(guess, e.goal)
}

// Humanize the result of a guess. Given a signature like "++##*"
// and guess like "after", the result is AF--r meaning the A and F
// are correcly placed, TE are not in the goal, and r is present
// but out of place.
func Humanize(signature string, guess string) string {
	var result strings.Builder
	for i, r := range signature {
		switch r {
		case LETTER_CORRECT:
			result.WriteRune(unicode.ToUpper(rune(guess[i])))
		case LETTER_IN_WORD:
			result.WriteRune(rune(guess[i]))
		case LETTER_WRONG:
			result.WriteRune('-')
		default:
			result.WriteRune('?')
			fmt.Fprintf(os.Stderr, "humanizing result string: invalid character %c in signature\n", r)
		}
	}
	return result.String()
}
