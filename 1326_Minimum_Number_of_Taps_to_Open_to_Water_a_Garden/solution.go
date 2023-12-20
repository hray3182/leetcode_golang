// 1326. Minimum Number of Taps to Open to Water a Garden

package leetcode

// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/

// 需要一個2維矩陣記錄灑水器的範圍
//
func minTaps(n int, ranges []int) int {
	// Find the reach of each position, NOT each tap. We do this by finding the left
	// and right ends of each tap, and then creating a list with the reach from the
	// left, aka pretending like that's where the water stream starts, and is only
	// shooting in a straight line.
	reach := make([]int, n+1)
	for i := 0; i < len(ranges); i++ {
		left := max(0, i-ranges[i])
		right := min(n, i+ranges[i])
		reach[left] = max(reach[left], right)
	}

	// Count how many taps it takes to get us to the max reach of all water streams.
	globalMax := reach[0]
	localMax := reach[0]
	taps := 1
	for i := 0; i < n; i++ {
		// If we've ever started checking a point that's farther than the farthest
		// reach that we know about, that means no water stream has been able to
		// reach this point yet, and no stream ever will (because we're only looking
		// at streams shooting in a straight line.)
		if i > globalMax {
			return -1
		}

		// Check if the reach from this position is farther than the absolute
		// maximum that we know about, and update if it is.
		globalMax = max(globalMax, reach[i])

		// If we've reached a position that's farther than the "current" maximum
		// that we're tracking, that means we need to be tracking a new water stream,
		// so we add 1 to our tap count, and set the local maximum that we're tracking
		// to the absolute maximum that we know about.
		if i == localMax {
			localMax = globalMax
			taps++
		}
	}
	return taps
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
