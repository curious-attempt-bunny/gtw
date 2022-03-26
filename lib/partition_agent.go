package gtw

import (
	"fmt"
	"math"
)

type PartitionScoreFn func(map[string]int) float64

func (agent *GtwAgent) GuessByPartitionScore(scoringMethod PartitionScoreFn) string {
	guessScores := make(map[string]float64)
	for _, candidateGuess := range agent.remaining {
		// partitionMap := PartitionWordCountsByGuessSignatures(agent.remaining, candidateGuess)
		partitionMap := PartitionWordCountsByGuessSignaturesCached(agent.remaining, candidateGuess)
		// fmt.Printf("guess %s has partition map %v\n", candidateGuess, partitionMap)
		guessScores[candidateGuess] = scoringMethod(partitionMap)
	}

	// fmt.Printf("%v\n", guessScores)

	bestGuess := "TODO"
	bestScore := float64(0)
	for guess, score := range guessScores {
		// break ties alphabetically for determinism
		if score > bestScore || (score == bestScore && guess < bestGuess) {
			// if score == bestScore {
			// 	print("Tie breaking ", guess, " vs ", bestGuess, "(", score, " and ", bestScore, ")\n")
			// } else {
			// 	print("Better ", guess, " vs ", bestGuess, "(", score, " and ", bestScore, ")\n")
			// }
			bestGuess = guess
			bestScore = score
		}
	}

	if bestGuess == "ridge" {
		fmt.Printf("Best is %s with partitions %v\n", bestGuess, PartitionWordCountsByGuessSignaturesCached(agent.remaining, bestGuess))
		fmt.Printf("Corpus is %v.\n", agent.remaining)
	}
	return bestGuess
}

func PartitionScoreByNumberOfPartitions(partitionMap map[string]int) float64 {
	return float64(len(partitionMap))
}

func PartitionScoreByMaximumPartitionSize(partitionMap map[string]int) float64 {
	maxSize := 0
	for _, size := range partitionMap {
		if size > maxSize {
			maxSize = size
		}
	}

	return float64(10000 - maxSize)
}

func PartitionScoreByPartitionSizeDeviation(partitionMap map[string]int) float64 {
	wordCount := 0
	for _, size := range partitionMap {
		wordCount += size
	}

	mean := float64(wordCount) / float64(len(partitionMap))

	sumDeviationSquared := 0.0
	for _, size := range partitionMap {
		deviation := float64(size) - mean
		sumDeviationSquared += deviation * deviation
	}

	standardDeviation := math.Sqrt(sumDeviationSquared)

	return 10000.0 - ((standardDeviation + 1) / float64(len(partitionMap)))
}
