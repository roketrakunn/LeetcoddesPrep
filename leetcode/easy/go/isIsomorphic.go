func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		cs := s[i]
		ct := t[i]

		if val, ok := sToT[cs]; ok && val != ct {
			return false
		}
		if val, ok := tToS[ct]; ok && val != cs {
			return false
		}

		sToT[cs] = ct
		tToS[ct] = cs
	}

	return true
}

