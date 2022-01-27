package main

func AlphaCount(s string) int {
	sentence := []byte(s)
	counter := 0
	for _, letter := range sentence {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			counter++
		}
	}
	return counter
}