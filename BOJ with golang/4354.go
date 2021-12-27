package main

import (
	"fmt"
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
	pi := getPi(p)
	var n int = len(s)
	var m int = len(p)
	var j int = 0
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		for j > 0 && s[i] != p[j] {
			j = pi[j-1]
		}
		if s[i] == p[j] {
			if j == m-1 {
				ans = append(ans, i-m+1)
				j = pi[j]
			} else {
				j += 1
			}
		}

	}

	return ans
}

func main() {

	for s := ""; s != "."; fmt.Scan(&s) {
		var n int = len(s)
		for i := 1; i <= n; i++ {
			var tmp string = s[0:i]
			ans := kmp(s, tmp)
			fmt.Println(ans)
		}
	}

}
