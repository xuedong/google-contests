package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main100() {
	var testCases int
	var B int
	fmt.Scanln(&testCases, &B)
	for testCase := 1; testCase <= testCases; testCase++ {
		var bits bytes.Buffer
		for i := 0; i < 150; i++ {
			fmt.Printf("%d\n", i%10+1)
			var rec int
			fmt.Scanln(&rec)
			if i >= 140 {
				bits.WriteString(strconv.Itoa(rec))
			}
		}

		output := bits.String()
		fmt.Printf("%s\n", output)
		var answer string
		fmt.Scanln(&answer)
		if answer == "N" {
			break
		}
	}
}
