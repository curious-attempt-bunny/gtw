package gtw

import (
	"reflect"
	"testing"
)

func TestInformRetainsWordsWithLettersInTheCorrectPlace(t *testing.T) {
	corpus := []string{"twist", "tabby", "tacky"}
	guess := "twist"
	answer := "tacky"
	signature, _ := ScoreAgainstGoal(guess, answer)

	nextAgent := Agent(corpus).Inform(guess, signature)

	// tabby and tease should be kept because the leading "t" is in the right place
	expectedRemaining := []string{"tabby", "tacky"}
	if !reflect.DeepEqual(nextAgent.remaining, expectedRemaining) {
		t.Errorf("Expected remaining corpus to be %v but was %v\n", expectedRemaining, nextAgent.remaining)
	}
}

func TestInformRemovesWordsWithGuessLettersNotInTheAnswer(t *testing.T) {
	corpus := []string{"twist", "tween", "tacky"}
	guess := "twist"
	answer := "tacky"
	signature, _ := ScoreAgainstGoal(guess, answer)

	nextAgent := Agent(corpus).Inform(guess, signature)

	// tween should be eliminated because "w" is not in "tacky"
	expectedRemaining := []string{"tacky"}
	if !reflect.DeepEqual(nextAgent.remaining, expectedRemaining) {
		t.Errorf("Expected remaining corpus to be %v but was %v\n", expectedRemaining, nextAgent.remaining)
	}
}

func TestInformRemovesWordsWithoutEnoughLettersInTheWrongPlace(t *testing.T) {
	corpus := []string{"twist", "table", "fatty"}
	guess := "twist"
	answer := "fatty"
	signature, _ := ScoreAgainstGoal(guess, answer)

	nextAgent := Agent(corpus).Inform(guess, signature)

	// table should be eliminated because it only has one "t" and the guess and answer have two
	expectedRemaining := []string{"fatty"}
	if !reflect.DeepEqual(nextAgent.remaining, expectedRemaining) {
		t.Errorf("Expected remaining corpus to be %v but was %v\n", expectedRemaining, nextAgent.remaining)
	}
}

func TestInform1(t *testing.T) {
	corpus := []string{"catty", "catch", "watch"}
	guess := "catty"
	answer := "watch"
	signature, _ := ScoreAgainstGoal(guess, answer)
	print(signature)

	nextAgent := Agent(corpus).Inform(guess, signature)

	// catch should be eliminated
	expectedRemaining := []string{"watch"}
	if !reflect.DeepEqual(nextAgent.remaining, expectedRemaining) {
		t.Errorf("Expected remaining corpus to be %v but was %v\n", expectedRemaining, nextAgent.remaining)
	}
}
