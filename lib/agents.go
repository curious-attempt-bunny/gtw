package gtw

func (agent *GtwAgent) GuessUsingMethod(method string) string {
	if method == "dumb" {
		return agent.remaining[0]
	} else if method == "partition_size" {
		return agent.GuessByPartitionScore(PartitionScoreByNumberOfPartitions)
	} else if method == "max_partition_size" {
		return agent.GuessByPartitionScore(PartitionScoreByMaximumPartitionSize)
	} else if method == "partition_size_deviation" {
		return agent.GuessByPartitionScore(PartitionScoreByPartitionSizeDeviation)
	} else {
		panic(method)
	}
}
