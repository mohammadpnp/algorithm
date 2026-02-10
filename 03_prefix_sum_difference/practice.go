package prefixsumdifference

func subarraysDivByK(nums []int, k int) int {
	count := 0
	prefixSum := 0
	freq := make(map[int]int)

	freq[0] = 1

	for _, num := range nums {
		prefixSum += num
		rem := prefixSum % k

		if rem < 0 {
			rem += k
		}

		count += freq[rem]
		freq[rem]++
	}

	return count
}
