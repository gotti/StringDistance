package main

import (
	"fmt"
	"math"
)

func GetJaroDistance(s1, s2 []rune) float64 {

	if len(s1) > len(s2) {
		var temp = s2
		s2 = s1
		s1 = temp
	}

	const (
		w1 = 0.33333333
		w2 = 0.33333333
		wt = 0.33333333
	)

	leng := int(math.Max(float64(len(s1)), float64(len(s2)))/2) - 1
	var c int
	var matched1 = make([]bool, len(s1))
	var matched2 = make([]bool, len(s2))

	for i := 0; i < len(s1); i++ {
		for j := -1 * leng; j <= leng; j++ {
			if i+j < 0 || int(len(s2)) <= i+j {
				continue
			}
			if s1[i] == s2[i+j] {
				c++
				matched1[i] = true
				matched2[i+j] = true
			}
		}
	}
	var sd1, sd2 []rune
	var numberOfReplace int
	for i := 0; i < len(s1); i++ {
		if matched1[i] {
			sd1 = append(sd1, s1[i])
		}
	}
	for i := 0; i < len(s2); i++ {
		if matched2[i] {
			sd2 = append(sd2, s2[i])
		}
	}
	fmt.Println(sd1)
	fmt.Println(sd2)
	for i := 0; i < int(math.Max(float64(len(sd1)), float64(len(sd2)))); i++ {
		if sd1[i] != sd2[i] {
			numberOfReplace++
		}
	}
	t := numberOfReplace / 2

	ret := w1*float64(c)/float64(len(s1)) + w2*float64(c)/float64(len(s2)) + wt*float64((c-t))/float64(c)
	return ret
}

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(GetJaroDistance([]rune(s1), []rune(s2)))
}
