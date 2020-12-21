package compare

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

var kingTime = reflect.TypeOf(time.Time{}).Kind()

// Compare 如下
//    k1 > k2 -->  1
//    k1 == k2 --> 0
//    k1 < k2 --> -1
type Compare func(k1, k2 interface{}) int

// RuneArray []rune compare
func RuneArray(k1, k2 interface{}) int {
	s1 := k1.([]rune)
	s2 := k2.([]rune)

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

// ByteArray []byte compare
func ByteArray(k1, k2 interface{}) int {
	s1 := k1.([]byte)
	s2 := k2.([]byte)

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

// String comp
func String(k1, k2 interface{}) int {
	s1 := k1.(string)
	s2 := k2.(string)

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

func Int(k1, k2 interface{}) int {
	c1 := k1.(int)
	c2 := k2.(int)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int8(k1, k2 interface{}) int {
	c1 := k1.(int8)
	c2 := k2.(int8)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int16(k1, k2 interface{}) int {
	c1 := k1.(int16)
	c2 := k2.(int16)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int32(k1, k2 interface{}) int {
	c1 := k1.(int32)
	c2 := k2.(int32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int64(k1, k2 interface{}) int {
	c1 := k1.(int64)
	c2 := k2.(int64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt(k1, k2 interface{}) int {
	c1 := k1.(uint)
	c2 := k2.(uint)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt8(k1, k2 interface{}) int {
	c1 := k1.(uint8)
	c2 := k2.(uint8)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt16(k1, k2 interface{}) int {
	c1 := k1.(uint16)
	c2 := k2.(uint16)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt32(k1, k2 interface{}) int {
	c1 := k1.(uint32)
	c2 := k2.(uint32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt64(k1, k2 interface{}) int {
	c1 := k1.(uint64)
	c2 := k2.(uint64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Float32(k1, k2 interface{}) int {
	c1 := k1.(float32)
	c2 := k2.(float32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Float64(k1, k2 interface{}) int {
	c1 := k1.(float64)
	c2 := k2.(float64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Byte(k1, k2 interface{}) int {
	c1 := k1.(byte)
	c2 := k2.(byte)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Rune(k1, k2 interface{}) int {
	c1 := k1.(rune)
	c2 := k2.(rune)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Time(k1, k2 interface{}) int {
	c1 := k1.(time.Time)
	c2 := k2.(time.Time)

	switch {
	case c1.After(c2):
		return 1
	case c1.Before(c2):
		return -1
	default:
		return 0
	}
}

// AutoComapre 通用比较. 自动判断. 效率对比其他低
func AutoComapre(k1, k2 interface{}) int {

	t1 := reflect.TypeOf(k1)
	t2 := reflect.TypeOf(k2)

	if t1.Kind() != t2.Kind() {
		panic("value1 value2 is not same type")
	}

	rv1 := reflect.ValueOf(k1)
	rv2 := reflect.ValueOf(k2)

	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
		t2 = t2.Elem()
		rv1 = rv1.Elem()
		rv2 = rv2.Elem()
	}

	switch t1.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v1 := rv1.Int()
		v2 := rv2.Int()
		switch {
		case v1 > v2:
			return 1
		case v1 < v2:
			return -1
		default:
			return 0
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v1 := rv1.Uint()
		v2 := rv2.Uint()
		switch {
		case v1 > v2:
			return 1
		case v1 < v2:
			return -1
		default:
			return 0
		}
	case reflect.Float32, reflect.Float64:
		v1 := rv1.Float()
		v2 := rv2.Float()
		switch {
		case v1 > v2:
			return 1
		case v1 < v2:
			return -1
		default:
			return 0
		}
	case reflect.String:
		v1 := rv1.String()
		v2 := rv2.String()
		return strings.Compare(v1, v2)
	case kingTime:
		v1 := rv1.Interface().(time.Time)
		v2 := rv1.Interface().(time.Time)
		switch {
		case v1.Before(v2):
			return 1
		case v1.After(v2):
			return -1
		default:
			return 0
		}
	default:

		panic(fmt.Sprintf("%v kind not handled", t1.Kind()))
	}

}
