package main

import "fmt"

func main() {
	var s1 string
	var n int
	fmt.Scanf("%d", &n)

	var cnt = 1
	var ans = 0
	for i := 0; i < n; i = i + 1 {
		fmt.Scanf("%s", &s1)
		ans = 0
		cnt = 0
		for j := 0; j < len(s1); j = j + 1 {
			if s1[j] == 'X' {
				cnt = 0
			} else if s1[j] == 'O' {
				cnt = cnt + 1
				ans = ans + cnt
			}
		}
		fmt.Println(ans)
	}
	fmt.Println()
}
