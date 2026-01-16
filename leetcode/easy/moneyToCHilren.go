package go_test

func distMoney(money int, children int) int {
	rem := money - children
	if rem < 0 {
		return -1
	}

	eights := rem / 7

	eights = min(eights , children )

	rem -= eights * 7

	if eights == children && rem > 0 {
		eights--
		return eights
	}

	if children-eights == 1 && rem == 3 {
		eights--
	}

	return eights
}

