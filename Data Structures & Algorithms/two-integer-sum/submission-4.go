func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))

	for i, num := range nums {
		hashmap[num] = i
	}

	for i, num := range nums {
		rem := target - num
		if id, ok := hashmap[rem]; ok && i != id {
			return []int{min(i, id), max(i, id)}
		}
	}

	return []int{}
}
