/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashtable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if p, ok := hashtable[target-nums[i]]; ok {
			return []int{p, i}
		}
		hashtable[nums[i]] = i
	}
	return nil
}

// @lc code=end

