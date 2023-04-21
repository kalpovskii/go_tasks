package main

import "fmt"

var romanToNun = map[rune]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func main() {
	romanStr := ""
	if _, err := fmt.Scan(&romanStr); err != nil {
		panic(err)
	}
	if len(romanStr) == 0 {
		fmt.Print(0)
		return
	}

	romanRunes := []rune(romanStr)
	res := 0
	for i := 0; i < len(romanRunes)-1; i++ {
		curr := romanToNun[romanRunes[i]]
		next := romanToNun[romanRunes[i+1]]
		if curr < next {
			res -= curr
		} else {
			res += curr
		}
	}
	res += romanToNun[romanRunes[len(romanRunes)-1]]
	fmt.Println(res)
}
