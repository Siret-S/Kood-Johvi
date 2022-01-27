package main

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	s0 := s
	if s[0] == 45 || s[0] == 43 {
		s = s[1:]
		if len(s) < 1 {
			return 0
		}
	}

	nm := 0
	for _, ch := range s {
		if !contain0to9(ch) {
			return 0
		}
		nm = nm*10 + charToInteg(ch)
	}

	if s0[0] == 45 {
		nm *= -1
	}
	return nm
}

func charToInteg(ch rune) int {
	count := 0
	for i := '1'; i <= ch; i++ {
		count++
	}
	return count
}

func contain0to9(ch rune) bool {
	for i := '0'; i <= '9'; i++ {
		if ch == i {
			return true
		}
	}
	return false
}
