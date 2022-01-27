package main

func AppendRange(min, max int) []int {
	if max <= min {
		return nil
	}
	var answer []int
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer
}
