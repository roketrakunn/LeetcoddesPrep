package go_test


func minimumDistance(nums []int) int {
	indexes := make(map[int][]int)

	const INF = int(1e9)
	minimum := INF

	for i, val := range nums {
		indexes[val] = append(indexes[val], i) // indices are already in increasing order
	}

	for _, v := range indexes {
		if len(v) < 3 {
			continue
		}
		for t := 0; t+2 < len(v); t++ {
			dist := 2 * (v[t+2] - v[t])
			if dist < minimum {
				minimum = dist
			}
		}
	}

	if minimum == INF {
		return -1
	}
	return minimum
}

