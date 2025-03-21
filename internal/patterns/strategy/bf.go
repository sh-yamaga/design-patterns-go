package strategy

type BruteForce struct{}

func (bf BruteForce) GuessNumber(target, max int) int {
	guess := 0
	for guess != target {
		guess++
	}
	return guess
}

func (bf BruteForce) Name() string {
	return "BruteForce"
}
