package go_test

func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := make(map[int]struct{})
    set2 := make(map[int]struct{})

    for _, x := range nums1 {
        set1[x] = struct{}{}
    }
    for _, x := range nums2 {
        set2[x] = struct{}{}
    }

    res := make([][]int, 2)

    // elements in nums1 but not nums2
    for x := range set1 {
        if _, ok := set2[x]; !ok {
            res[0] = append(res[0], x)
        }
    }

    // elements in nums2 but not nums1
    for x := range set2 {
        if _, ok := set1[x]; !ok {
            res[1] = append(res[1], x)
        }
    }

    return res
}

