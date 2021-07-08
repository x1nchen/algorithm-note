package main

import (
	"fmt"
	"math"
)

func DP(w []int, v []int, N, W int) int {
	// dp[tn][rw] 其中 tn 表示第几个物品，rw 表示背包还能容纳的重量
	dp := make([][]int, N+1)

	// init
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for tn := 1; tn < N+1; tn++ { // 遍历每一件物品
		for rw := 1; rw < W+1; rw++ { // 背包容量有多大就还要计算多少次
			if rw < w[tn] {
				// 当背包容量小于第tn件物品重量时，只能放入前tn-1件
				dp[tn][rw] = dp[tn-1][rw]
				fmt.Printf("dp[%d][%d]=%d\n",tn,rw, dp[tn][rw])
			} else {
				// 当背包容量还大于第tn件物品重量时，进一步作出决策
				// 不放入的价值
				nValue := dp[tn-1][rw]
				// 放入的价值
				pValue := dp[tn-1][rw-w[tn]] + v[tn]
				dp[tn][rw] = int(math.Max(float64(nValue), float64(pValue)))
				fmt.Printf("dp[%d][%d]=%d\n",tn,rw, dp[tn][rw])
			}
		}
	}
	return dp[N][W]
}

// 0-1 背包问题
func main() {
	N, W := 3, 5             // 物品的总数，背包能容纳的总重量
	w := []int{0, 3, 2, 1}   // 物品的重量
	v := []int{0, 5, 2, 3}   // 物品的价值
	result := DP(w, v, N, W) // 输出答案
	fmt.Println(result)
}
