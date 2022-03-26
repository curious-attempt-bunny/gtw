package gtw

import "testing"

func TestScoreByMaximumNumberOfPartitions(t *testing.T) {
	corpus := []string{"twist", "tabby", "tease", "tweep", "twonk"}
	guess := Agent(corpus).GuessByPartitionScore(PartitionScoreByNumberOfPartitions)

	expected := "tease"
	if guess != expected {
		t.Errorf("Expected %s. Got %s. %s should have the greatest number of partitions.\n", expected, guess, expected)
	}
}

func TestScoreByMaximumPartitionSize(t *testing.T) {
	corpus := []string{"twist", "tabby", "tease", "tweep", "twonk"}
	guess := Agent(corpus).GuessByPartitionScore(PartitionScoreByMaximumPartitionSize)

	expected := "tease"
	if guess != expected {
		t.Errorf("Expected %s. Got %s. %s should have a max partition size of 1 (the smallest for all words).\n", expected, guess, expected)
	}
}
