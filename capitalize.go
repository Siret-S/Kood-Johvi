package main

func Capitalize(s string) string {
	res := []rune(s) // k6igepealt k6ik vaikesteks tahtedeks
	for index, x := range res {
		if x >= 'A' && x <= 'Z' {
			res[index] = x + 32
		}
	}
	for i, v := range res {
		if v >= 'a' && v <= 'z' && i != 0 {
			if (res[i-1] < 'A' || res[i-1] > 'Z') && (res[i-1] < 'a' || res[i-1] > 'z') && (res[i-1] < '0' || res[i-1] > '9') { // if it's a space or some other character other than letters or numbers
				res[i] = v - 32
			}
		} else if i == 0 && v >= 'a' && v <= 'z' {
			res[i] = v - 32
		}
	}
	return string(res)
}
