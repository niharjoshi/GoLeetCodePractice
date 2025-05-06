func twoSum(nums []int, target int) []int {

	visited := make(map[int]int)

	for pos, val := range nums {

		diff := target - val

		if p, ok := visited[diff]; ok {
			return []int{p, pos}
		}

		visited[val] = pos
	}

	return nil

}