package gtw

import (
	"os"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	partitions := PartitionWordCountsByGuessSignatures([]string{"twist", "total", "tuple", "tabby"}, "twist")

	count, exists := partitions["+++++"]
	if count != 1 {
		t.Errorf("Expected entry +++++ to have count of 1. Was %d, %t\n", count, exists)
	}

	count, exists = partitions["+###*"]
	if count != 1 {
		t.Errorf("Expected entry +###* to have count of 1. Was %d, %t\n", count, exists)
	}

	count, exists = partitions["+####"]
	if count != 2 {
		t.Errorf("Expected entry +#### to have count of 2. Was %d, %t\n", count, exists)
	}
}

func TestPartitionCache(t *testing.T) {
	corpus := []string{"twist", "total", "tuple", "tabby"}
	guess := "twist"

	cacheFilename := CacheFilename(corpus, guess)
	defer os.Remove(cacheFilename)

	partitionsUncached := PartitionWordCountsByGuessSignatures(corpus, guess)
	partitionsCached := PartitionWordCountsByGuessSignaturesCached(corpus, guess)

	if !reflect.DeepEqual(partitionsCached, partitionsUncached) {
		t.Errorf("Expected partitions to match: %v vs %v\n", partitionsCached, partitionsUncached)
	}

	_, err := os.Stat(cacheFilename)
	if err != nil {
		t.Errorf("Expected %s to exist: %v\n", cacheFilename, err)
	}

	partitionsCached = PartitionWordCountsByGuessSignaturesCached(corpus, guess)

	if !reflect.DeepEqual(partitionsCached, partitionsUncached) {
		t.Errorf("Expected cached retrieved partitions to match: %v vs %v\n", partitionsCached, partitionsUncached)
	}
}
