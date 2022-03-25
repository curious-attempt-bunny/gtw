package agent

import (
	// "fmt"
	// "strings"
	"testing"

	"github.com/gmofishsauce/gtw/lib"
)

func TestBenchmarkPartitionGuess(t *testing.T) {
	corpus, _ := gtw.LoadFile("../cmd/cli/webster-2-all-five-letter.corpus")
	
	answer := "depot"
	round := 0
	for len(corpus) > 1 {
		round = round + 1
		engine := gtw.New(corpus)
		engine.NewFixedGame(answer)
		a := Agent(corpus)
		
		guess := a.PartitionMaxGuess()
		signature, _ := engine.Score(guess)
		a = a.Inform(guess, signature)
		
		// t.Errorf("Round %d guess of %s for %s yields %s (%d words -> %d words)\n",
		// 	round,
		// 	guess,
		// 	answer,
		// 	signature,
		// 	len(corpus),
		// 	len(a.remaining))
		
		corpus = a.remaining
	}

	if round != 1000 {
		t.Errorf("Expected PartitionGuess to take 1000 rounds. Instead it took %d\n", round)
	}
}