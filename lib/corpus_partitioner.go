package gtw

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
)

func PartitionWordCountsByGuessSignaturesCached(corpus []string, guess string) map[string]int {
	cacheFilename := CacheFilename(corpus, guess)

	content, err := ioutil.ReadFile(cacheFilename)
	if err != nil {
		print("Calculating uncached partition counts for guess ", guess, "\n")
		data := PartitionWordCountsByGuessSignatures(corpus, guess)
		content, err = json.Marshal(data)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(cacheFilename, content, fs.ModePerm)
	}

	var result map[string]int
	json.Unmarshal(content, &result)

	return result
}

func CacheFilename(corpus []string, guess string) string {
	hashData := ""
	for _, word := range corpus {
		hashData = hashData + word
	}
	return fmt.Sprintf("../data/%x", sha256.Sum256([]byte(hashData)))
}

func PartitionWordCountsByGuessSignatures(corpus []string, guess string) map[string]int {
	partionMap := make(map[string]int)

	for _, word := range corpus {
		signature, _ := ScoreAgainstGoal(guess, word)

		partition, exists := partionMap[signature]

		if !exists {
			partition = 0
		}

		partionMap[signature] = partition + 1
	}

	return partionMap
}
