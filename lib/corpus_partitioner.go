package gtw

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"reflect"
)

func PartitionWordCountsByGuessSignaturesCached(corpus []string, guess string) map[string]int {
	cacheFilename := CacheFilename(corpus, guess)

	content, err := ioutil.ReadFile(cacheFilename)
	if err != nil {
		// print("Calculating uncached partition counts for guess ", guess, " (remaining count ", len(corpus), ", cache file ", cacheFilename, ")\n")
		data := PartitionWordCountsByGuessSignatures(corpus, guess)
		content, err = json.Marshal(data)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(cacheFilename, content, fs.ModePerm)
		if err != nil {
			panic(err)
		}

		var result map[string]int
		json.Unmarshal(content, &result)
		if !reflect.DeepEqual(result, data) {
			panic("Not equal!!")
		}
	}

	var result map[string]int
	json.Unmarshal(content, &result)

	return result
}

func CacheFilename(corpus []string, guess string) string {
	hash := sha256.New()
	for _, word := range corpus {
		io.WriteString(hash, word)
	}
	io.WriteString(hash, guess)
	return fmt.Sprintf("data/%x-%d.json", hash.Sum(nil), len(corpus))
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
		// fmt.Printf("%s against %s -> %s. Current count of %d\n", guess, word, signature, partition+1)
	}

	return partionMap
}
