func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    for i, num := range nums {
        if j, exists := dict[target - num]; exists {
            return []int{j, i}
        }
        dict[num] = i
    }

    return []int{}
}

