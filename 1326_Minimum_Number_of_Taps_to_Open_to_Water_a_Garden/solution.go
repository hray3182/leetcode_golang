// 1326. Minimum Number of Taps to Open to Water a Garden

package leetcode

import "sort"

// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/

// 需要一個2維矩陣記錄灑水器的範圍
//
func minTaps(n int, ranges []int) int {
	taps := make([][]int, n+1)
	for i := range taps {
		taps[i] = make([]int, 2)
		// 記錄灑水器範圍
		taps[i][0] = max(0, i-ranges[i]) // 左邊界
		taps[i][1] = min(n, i+ranges[i]) // 右邊界
	}
	// 排序矩陣
	// 規則如下:
	// 1. 左邊界小的在前面
	// 2. 左邊界相同的, 右邊界大的在前面
	sort.Slice(taps, func(i, j int) bool {
		if taps[i][0] == taps[j][0] {
			return taps[i][1] > taps[j][1]
		}
		return taps[i][0] < taps[j][0]
	})
	// 檢查第一個灑水器有沒有覆蓋到最左邊(0)
	if taps[0][0] > 0 {
		return -1

	}
	// 記錄右邊界
	right := taps[0][1]
	// 記錄水龍頭數量
	count := 0

	for i := range taps {
		count++
		// 如果 right >= n 的話，表示已經覆蓋整個花園，此時就可以結束
		if right >= n {
			return count
		}
		currRight := right
		for j := i + 1; j < len(taps); j++ {
			// 如果 j 的覆蓋範圍比 i 還大的話，則將 i 取代為 j
			if taps[j][0] <= right && taps[j][1] >= currRight {
				i = j
				currRight = taps[j][1]
			}
		}
		// 如果 j 遍歷完沒有超過一開始記錄的 right 的話，則代表一開始的範圍已經是最大的了
		// 接下來無論如何都無法覆蓋整個花園
		if currRight == right {
			return -1
		}
		right = currRight

		// 由於 i 會自增，所以需要減 1，用當前找到最大範圍的 i 來做為下一次循環的開始
		i--
	}
	return -1
}

// 從別人的解答看到的，超級優雅解法
func minTaps2(n int, ranges []int) int {
	reach := make([]int, n+1)
	for i := range ranges {
		// 用 left 當作 index, right 做為value
		left := max(0, i-ranges[i])
		right := min(n, i+ranges[i])
		reach[left] = max(reach[left], right)
	}

	if n > 0 && reach[0] == 0 {
		return -1
	}
	globalMax, localMax := reach[0], reach[0]
	count := 1
	for i := 1; i <= n; i++ {
		if i > globalMax {
			return -1
		}
		if i > localMax {
			localMax = globalMax
			count++
		}
		globalMax = max(globalMax, reach[i])
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
