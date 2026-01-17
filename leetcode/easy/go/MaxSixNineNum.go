package go_test

import "strconv"
func maximum69Number (num int) int {
	s := []byte(strconv.Itoa(num))
	for i := range len(s) { 
		if s[i] == '6' { 
			s[i] = '9'
			break
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
	
}
