package strategy

type BinarySearch struct{}

func (bs BinarySearch) GuessNumber(target, max int) int {
	min := 0
	for min <= max {
		guess := (min + max) / 2
		if guess < target {
			min = guess + 1
		} else if target < guess {
			max = guess - 1
		} else {
			return guess
		}
	}

	// Return -1 if the target is not found
	return -1
}

func (bs BinarySearch) Name() string {
	return "BinarySearch"
}
