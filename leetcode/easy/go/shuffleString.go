package go_test


func restoreString(s string, indices []int) string {
	n := len(s)
	results := make([]byte , n)

	for i , ch := range s {
		ind := indices[i]
		results[ind]  = byte(ch)
	}
	return string(results)	
}
