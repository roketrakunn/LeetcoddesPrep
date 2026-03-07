package go_test

func repeatedSubstringPattern(s string) bool {
      n := len(s)
      for l := 1; l <= n/2; l++ {
          if n%l != 0 {
              continue
          }
          sub := s[:l]
          match := true
          for i := l; i < n; i += l {
              if s[i:i+l] != sub {
                  match = false
                  break
              }
          }
          if match {
              return true
          }
      }
    
	return false
}

