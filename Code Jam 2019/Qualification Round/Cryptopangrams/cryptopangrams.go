package main

import (
	"bytes"
	"fmt"
	"sort"
)

func mapkey(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func unique(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// func gcd(u, v uint64) uint64 {
// 	var shift uint
//
// 	/* GCD(0,v) == v; GCD(u,0) == u, GCD(0,0) == 0 */
// 	if u == 0 {
// 		return v
// 	}
// 	if v == 0 {
// 		return u
// 	}
//
// 	/* Let shift := lg K, where K is the greatest power of 2
// 	   dividing both u and v. */
// 	for shift = 0; ((u | v) & 1) == 0; shift++ {
// 		u >>= 1
// 		v >>= 1
// 	}
//
// 	// While u is even, continue dividing by 2 (right-shift) until it becomes odd
// 	for (u & 1) == 0 {
// 		u >>= 1
// 	}
//
// 	/* From here on, u is always odd. */
// 	for ok := true; ok; ok = (v != 0) {
// 		/* remove all factors of 2 in v -- they are not common */
// 		/*   note: v is not zero, so while will terminate */
// 		for (v & 1) == 0 /* Loop X */ {
// 			v >>= 1
// 		}
//
// 		/* Now u and v are both odd. Swap if necessary so u <= v,
// 		   then set v = v - u (which is even). For bignums, the
// 		   swapping is just pointer movement, and the subtraction
// 		   can be done in-place. */
// 		if u > v {
// 			// Swap u and v.
// 			//  unsigned int t = v; v = u; u = t;}
// 			v, u = u, v
// 		}
// 		v = v - u // Here v >= u.
// 	}
//
// 	/* restore common factors of 2 */
// 	return u << shift
// }

func gcd(a int, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	return a
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var limit int
		var length int
		fmt.Scanln(&limit, &length)

		list := make([]int, length)
		for i := range list {
			fmt.Scan(&list[i])
		}

		gcdList := make([]int, length+1)
		for i := 1; i < length; i++ {
			gcdList[i] = gcd(list[i-1], list[i])
		}

		begin := list[0] / gcdList[1]
		end := list[length-1] / gcdList[length-1]
		gcdList[0] = begin
		gcdList[length] = end
		// fmt.Printf("%v\n", gcdList)

		// uniqueList := unique(gcdList)
		// copy := make([]int, 26)
		// for i := 0; i < 26; i++ {
		// 	copy[i] = int(uniqueList[i])
		// }
		// sort.Ints(copy)
		// uniqueListCopy := make([]uint64, 26)
		// for i := 0; i < 26; i++ {
		// 	uniqueListCopy[i] = uint64(copy[i])
		// }
		copy := make([]int, length+1)
		for i := 0; i < length+1; i++ {
			copy[i] = gcdList[i]
		}
		sort.Ints(copy)
		uniqueList := unique(copy)
		// fmt.Printf("%v\n", uniqueList)

		var m map[string]int
		m = make(map[string]int)
		for i := 0; i < 26; i++ {
			m[string(i+65)] = uniqueList[i]
		}
		// fmt.Printf("%v\n", m)

		var buffer bytes.Buffer
		for i := 0; i < length+1; i++ {
			key, _ := mapkey(m, gcdList[i])
			buffer.Write([]byte(key))
		}
		// fmt.Printf("%v\n", gcdList)

		var result string
		result = buffer.String()
		fmt.Printf("Case #%d: %s\n", testCase, result)
	}
}
