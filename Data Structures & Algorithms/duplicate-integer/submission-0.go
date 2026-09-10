func hasDuplicate(nums []int) bool {
    hashset := map[int]struct{}{}    

    for _, num := range nums {
        if _, ok := hashset[num]; ok {
            return true
        }
        hashset[num] = struct{}{}
    }

    return false
}
