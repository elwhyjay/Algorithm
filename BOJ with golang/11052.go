package main

import (
	"fmt"
)

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int64, n+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i+1])
	}
	dp := make([]int64, n+1)
	dp[1] = arr[1]
	for i := 2; i <= n; i++ {

		for j := 1; j <= i; j++ {

			dp[i] = Max(dp[i], dp[i-j]+dp[j])
		}
		dp[i] = Max(dp[i], arr[i])
	}
	fmt.Print(dp[n])
}
