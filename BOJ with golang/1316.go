package main

import "fmt"

func main() {
	var s string
	var n int
	var ans int = 0
	alphabet := [26]int{}
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i += 1 {
		var flag bool = false
		for j := 0; j < 26; j++ {
			alphabet[j] = 0
		}
		fmt.Scanf("%s", &s)
		alphabet[s[0]-97] = 1
		for j := 1; j < len(s); j++ {
			if alphabet[s[j]-97] == 1 && s[j-1] != s[j] {
				flag = true
			}
			alphabet[s[j]-97] = 1

		}

		if flag == true {
			continue
		} else {
			ans++
		}
	}
	fmt.Print(ans)

}
