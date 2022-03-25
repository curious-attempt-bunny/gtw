package agent

import (
	"math"

	"github.com/gmofishsauce/gtw/lib"
)

func (agent *GtwAgent) PartitionGuess() string {
	best_partition_score := 0
	best_guess := "nil" // TODO
	for _, guess := range(agent.remaining) {
		// print("Considering ", guess, "\n")
		partitionMap := PartitionByGuessSignatures(agent.remaining, guess)
		// print("Guess ", guess, " has ", len(partitionMap), " partitions\n")

		score := len(partitionMap)
		if score > best_partition_score {
			best_partition_score = score
			best_guess = guess
		}
	}

	print("Guess ", best_guess, " has ", best_partition_score, " partitions (best).\n")
	return best_guess
}

func (agent *GtwAgent) PartitionMaxGuess() string {
	best_partition_score := 0
	best_guess := "nil" // TODO
	for _, guess := range(agent.remaining) {
		// print("Considering ", guess, "\n")
		partitionMap := PartitionByGuessSignatures(agent.remaining, guess)
		// print("Guess ", guess, " has ", len(partitionMap), " partitions\n")

		score := 0
		for _, partition := range(partitionMap) {
			// print("Guess ", guess, " has partition with signature ", signature, " and ", len(partition), " words\n")
			if len(partition) > score {
				score = len(partition)
			}
		}

		// print("^^ score is ", score)
		if best_guess == "nil" || score < best_partition_score {
			best_partition_score = score
			best_guess = guess
			// print("^^ best so far\n")
		}
	}

	print("Guess ", best_guess, " has biggest partition of size ", best_partition_score, " words (best).\n")
	return best_guess
}

func (agent *GtwAgent) PartitionDeviationGuess() string {
	best_partition_score := 0.0
	best_guess := "nil" // TODO
	for _, guess := range(agent.remaining) {
		// print("Considering ", guess, "\n")
		partitionMap := PartitionByGuessSignatures(agent.remaining, guess)
		// print("Guess ", guess, " has ", len(partitionMap), " partitions\n")

		total := 0
		
		for _, partition := range(partitionMap) {
			// print("Guess ", guess, " has partition with signature ", signature, " and ", len(partition), " words\n")
			total += len(partition)
		}
		mean := float64(total) / float64(len(partitionMap))

		std_squared := 0.0
		for _, partition := range(partitionMap) {
			// print("Guess ", guess, " has partition with signature ", signature, " and ", len(partition), " words\n")
			deviation := float64(len(partition)) - mean
			std_squared += deviation*deviation
		}
		std := math.Sqrt(std_squared)

		score := (std + 1.0) / (float64(len(partitionMap)) + 1.0)

		// print("^^ score is ", score)
		if best_guess == "nil" || score < best_partition_score {
			best_partition_score = score
			best_guess = guess
			// print("^^ best so far\n")
		}
	}

	print("Guess ", best_guess, " has smoothed standard deviation of ", best_partition_score, " (best).\n")
	return best_guess
}

func (agent *GtwAgent) PartitionAverageSizeGuess() string {
	best_partition_score := 0.0
	best_guess := "nil" // TODO
	for _, guess := range(agent.remaining) {
		// print("Considering ", guess, "\n")
		partitionMap := PartitionByGuessSignatures(agent.remaining, guess)
		// print("Guess ", guess, " has ", len(partitionMap), " partitions\n")

		total := 0
		
		for _, partition := range(partitionMap) {
			// print("Guess ", guess, " has partition with signature ", signature, " and ", len(partition), " words\n")
			total += len(partition)
		}
		mean := float64(total) / float64(len(partitionMap))
		score := mean
		
		// print("^^ score is ", score)
		if best_guess == "nil" || score < best_partition_score {
			best_partition_score = score
			best_guess = guess
			// print("^^ best so far\n")
		}
	}

	print("Guess ", best_guess, " has average partition size of ", best_partition_score, " (best).\n")
	return best_guess
}

func PartitionByGuessSignatures(corpus []string, guess string) map[string][]string {
	partionMap := make(map[string][]string)

	for _, word := range(corpus) {
		signature, _ := gtw.ScoreAgainstGoal(guess, word)

		partition, exists := partionMap[signature]

		if !exists {
			partition = make([]string, 1)
		}

		partionMap[signature] = append(partition, word)
	}

	return partionMap
}

