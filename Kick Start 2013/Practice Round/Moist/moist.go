package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strconv"
	"strings"
)

type MyInput struct {
	rdr io.Reader
	lineChan chan string
	initialized bool
}

func (mi *MyInput) start(done chan struct{}) {
	r := bufio.NewReader(mi.rdr)
	defer func() { close(mi.lineChan) }()
	for {
		line, err := r.ReadString('\n')
		if !mi.initialized {
			mi.initialized = true
			done <- struct{}{}
		}
		mi.lineChan <- strings.TrimSpace(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}

func (mi *MyInput) readLine() string {
	// if this is the first call, initialize
	if !mi.initialized {
		mi.lineChan = make(chan string)
		done := make(chan struct{})
		go mi.start(done)
		<-done
	}

	res, ok := <-mi.lineChan
	if !ok {
		panic("trying to read from a closed channel")
	}
	return res
}

func (mi *MyInput) readInt() int {
	line := mi.readLine()
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *MyInput) readInt64() int64 {
	line := mi.readLine()
	i, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (mi *MyInput) readInts() []int {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int{}
	for _, s := range parts {
		tmp, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *MyInput) readInt64s() []int64 {
	line := mi.readLine()
	parts := strings.Split(line, " ")
	res := []int64{}
	for _, s := range parts {
		tmp, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, tmp)
	}
	return res
}

func (mi *MyInput) readWords() []string {
	line := mi.readLine()
	return strings.Split(line, " ")
}

func strcmp(s1, s2 string) int {
	l1, l2 := len(s1), len(s2)
	var length int
	if l1 > l2 {
		length = l2
	} else {
		length = l1
	}

	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
	}
	return l1 - l2
}

func main() {
    mi := MyInput{rdr: os.Stdin}
	testCases := mi.readInt()

	for testCase := 1; testCase <= testCases; testCase++ {
		n := mi.readInt()

		ans := 0
		maxString := mi.readLine()
		for i := 1; i <= n-1; i++ {
			name := mi.readLine()
			if strcmp(name, maxString) < 0 {
				ans++
			} else {
			    maxString = name
			}
		}

		fmt.Printf("Case #%d: %d\n", testCase, ans)
	}
}
