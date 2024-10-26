//go:build bucketsort
// +build bucketsort

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

package medium

// @lc code=start

func topKFrequent(nums []int, k int) (result []int) {
	// count the number
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// create bucket from the freq
	// bucket size maximum is the length of array
	// add +1 to start from index 1
	bucket := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		bucket[freq] = append(bucket[freq], num)
	}

	// loop from the end of bucket
	// because the most freqs is over there
	for i := len(bucket) - 1; i > 0; i-- {
		// add bucket value to result with limit of k
		// no need extra check because result is guaranteed to be unique
		if len(result) < k {
			result = append(result, bucket[i]...)
		}
	}

	return
}

// @lc code=end
