package hashmapsetfrequency

func checkFrequentNumberAndisEqualToK(nums []int, k int) bool {
	freq := make(map[int]int)

	for i, x := range nums {
		if j, ok := freq[x]; ok {

			if i-j <= k {
				return true
			}
		}
		freq[x] = i
	}
	return false
}
