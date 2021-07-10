package listtree

type Compare func(interface{}, interface{}) int

// CompatorByte 默认比较字节
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

// CompatorMath 字节看书法的差异, 使用与数学类比较
func CompatorMath(i1, i2 interface{}) int {

	s1 := i1.([]byte)
	s2 := i2.([]byte)

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
