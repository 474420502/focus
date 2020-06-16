package vtree

func CompatorByte(s1, s2 []byte) int {

	switch {
	case len(s1) > len(s2):
		for i := 0; i < len(s2); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 1
	case len(s1) < len(s2):
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return -1
	default:
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}

func CompatorMath(s1, s2 []byte) int {

	switch {
	case len(s1) > len(s2):
		// for i := 0; i < len(s2); i++ {
		// 	if s1[i] != s2[i] {
		// 		if s1[i] > s2[i] {
		// 			return 1
		// 		}
		// 		return -1
		// 	}
		// }
		return 1
	case len(s1) < len(s2):
		// for i := 0; i < len(s1); i++ {
		// 	if s1[i] != s2[i] {
		// 		if s1[i] > s2[i] {
		// 			return 1
		// 		}
		// 		return -1
		// 	}
		// }
		return -1
	default:
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}
