package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getPi(p string) []int {

	var n int = len(p)
	var j int = 0
	arr := make([]int, n)
	for i := 1; i < n; i++ {
		for j > 0 && p[i] != p[j] {
			j = arr[j-1]
		}
		if p[i] == p[j] {
			j++
			arr[i] = j
		}
	}
	return arr
}

func kmp(s string, p string) []int {
	ans := make([]int, 0)
	pi := getPi(p)
	var n int = len(s)
	var m int = len(p)
	var begin int = 0
	var j int = 0
	for begin <= n-m {
		if j < m && s[begin+j] == p[j] {
			j++

			if j == m {
				ans = append(ans, begin)
			}
		} else {
			if j == 0 {
				begin++
			} else {
				begin += j - pi[j-1]
				j = pi[j-1]
			}
		}
	}

	return ans
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	p, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	p = strings.TrimSuffix(p, "\n")
	a := kmp(s, p)
	fmt.Println(len(a))
	for _, val := range a {
		fmt.Println(val + 1)
	}
}
