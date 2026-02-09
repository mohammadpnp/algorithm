package twopointersslidingwindow

func longestSubArrayOf1(num string, k int) int {
	left, max, zeroCount := 0, 0, 0

	for right := 0; right < len(num); right++ {
		if num[right] == '0' {
			zeroCount++
		}

		for zeroCount > k {
			if num[left] == '0' {
				zeroCount--
			}
			left++
		}

		if right-left+1 > max {
			max = right - left + 1
		}
	}

	return max

}
