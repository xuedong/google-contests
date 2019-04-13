package main

import (
	"fmt"
	"sort"
)

// MinMax returns the indices of the smallest and largest element of an int slice
func MinMax(array []int) (int, int, int, int) {
	var max = array[0]
	var min = array[0]
	maxInd := 0
	minInd := 0
	for index, value := range array {
		if max < value {
			max = value
			maxInd = index
		}
		if min > value {
			min = value
			minInd = index
		}
	}
	return minInd, maxInd, min, max
}

// LCP returns the longest common xfix of a set of strings
func LCP(l []string) string {
	// Special cases first
	switch len(l) {
	case 0:
		return ""
	case 1:
		return l[0]
	}
	// LCP of min and max (lexigraphically)
	// is the LCP of the whole set.
	min, max := l[0], l[0]
	for _, s := range l[1:] {
		switch {
		case s < min:
			min = s
		case s > max:
			max = s
		}
	}
	for i := 0; i < len(min) && i < len(max); i++ {
		if min[i] != max[i] {
			return min[:i]
		}
	}
	// In the case where lengths are not equal but all bytes
	// are equal, min is the answer ("foo" < "foobar").
	return min
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var numberWords int
		fmt.Scanln(&numberWords)

		strList := make([]string, numberWords)

		for i := 1; i <= numberWords; i++ {
			var instr string
			fmt.Scanln(&instr)
			strList[i-1] = Reverse(instr)
		}

		sort.Strings(strList)
		length := len(strList)
		rhymes := 0
		currentLCP := ""

		for length > 1 {
			if length == 2 {
				lcp := LCP([]string{strList[0], strList[1]})
				if len(lcp) > 0 && currentLCP != LCP([]string{strList[0], strList[1]}) {
					rhymes += 2
				}
				length--
			} else {
				prefixLengths := make([]int, length-1)
				for i := 0; i < length-1; i++ {
					lcp := LCP([]string{strList[i], strList[i+1]})
					prefixLengths[i] = len(lcp)
				}
				_, maxInd, _, max := MinMax(prefixLengths)

				if max > 0 && currentLCP != LCP([]string{strList[maxInd], strList[maxInd+1]}) {
					rhymes += 2
				}

				currentLCP = LCP([]string{strList[maxInd], strList[maxInd+1]})

				copy(strList[maxInd:], strList[maxInd+2:])
				strList[len(strList)-1] = ""
				strList[len(strList)-2] = ""
				strList = strList[:len(strList)-1]
				length = len(strList)
			}
		}

		fmt.Printf("Case #%d: %d\n", testCase, rhymes)
	}
}
