package tried

var WordIndexDict map[WordIndexType]*wordIndexStore

func init() {
	WordIndexDict = make(map[WordIndexType]*wordIndexStore)
	WordIndexDict[WordIndexLower] = &wordIndexStore{WordIndexLower, wordIndexLower, indexWordLower, 26}
	WordIndexDict[WordIndexUpper] = &wordIndexStore{WordIndexUpper, wordIndexUpper, indexWordUpper, 26}
	WordIndexDict[WordIndexDigital] = &wordIndexStore{WordIndexDigital, wordIndexDigital, indexWordDigital, 10}
	WordIndexDict[WordIndexUpperLower] = &wordIndexStore{WordIndexUpperLower, wordIndexUpperLower, indexWordUpperLower, 52}
	WordIndexDict[WordIndexLowerDigital] = &wordIndexStore{WordIndexLowerDigital, wordIndexLowerDigital, indexWordLowerDigital, 36}
	WordIndexDict[WordIndexUpperDigital] = &wordIndexStore{WordIndexUpperDigital, wordIndexUpperDigital, indexWordUpperDigital, 36}
	WordIndexDict[WordIndexUpperLowerDigital] = &wordIndexStore{WordIndexUpperLowerDigital, wordIndexUpperLowerDigital, indexWordUpperLowerDigital, 62}
	WordIndexDict[WordIndex256] = &wordIndexStore{WordIndex256, wordIndex256, indexWord256, 256}
	WordIndexDict[WordIndex32to126] = &wordIndexStore{WordIndex32to126, wordIndex32to126, indexWord32to126, ('~' - ' ' + 1)}
}

// WordIndexType 单词统计的类型 eg. WordIndexLower 意味Put的单词只支持小写...
type WordIndexType int

const (
	_ WordIndexType = iota
	WordIndexLower
	WordIndexUpper
	WordIndexDigital
	WordIndexUpperLower
	WordIndexLowerDigital
	WordIndexUpperDigital
	WordIndexUpperLowerDigital
	WordIndex256
	WordIndex32to126
)

type wordIndexStore struct {
	Type       WordIndexType
	Byte2Index func(byte) uint
	Index2Byte func(uint) byte
	DataSize   uint
}

func wordIndexLower(w byte) uint {
	return uint(w) - 'a'
}

func indexWordLower(w uint) byte {
	return byte(w) + 'a'
}

//
func wordIndexUpper(w byte) uint {
	return uint(w) - 'A'
}

func indexWordUpper(w uint) byte {
	return byte(w) + 'A'
}

//
func wordIndexDigital(w byte) uint {
	return uint(w) - '0'
}

func indexWordDigital(w uint) byte {
	return byte(w) + '0'
}

//
func wordIndexUpperLower(w byte) uint {
	iw := uint(w)
	if iw >= 'a' {
		return iw - 'a'
	}
	return iw - 'A' + 26
}

func indexWordUpperLower(w uint) byte {

	if w >= 26 {
		return byte(w) + 'A'
	}
	return byte(w) + 'a'
}

//
func wordIndexLowerDigital(w byte) uint {
	iw := uint(w)
	if iw >= 'a' {
		return iw - 'a'
	}
	return iw - '0' + 26
}

func indexWordLowerDigital(w uint) byte {
	if w >= 26 {
		return byte(w) + '0'
	}
	return byte(w) + 'a'
}

//
func wordIndexUpperDigital(w byte) uint {
	iw := uint(w)
	if iw >= 'A' {
		return iw - 'A'
	}
	return iw - '0' + 26
}

func indexWordUpperDigital(w uint) byte {
	if w >= 26 {
		return byte(w) + '0'
	}
	return byte(w) + 'a'
}

//
func wordIndexUpperLowerDigital(w byte) uint {
	iw := uint(w)
	if iw >= 'a' {
		return iw - 'a'
	} else if iw >= 'A' {
		return iw - 'A' + 26
	}
	return iw - '0' + 52
}

func indexWordUpperLowerDigital(w uint) byte {
	if w >= 52 {
		return byte(w) + '0'
	} else if w >= 26 {
		return byte(w) + 'A'
	}
	return byte(w) + 'a'
}

// wordIndex256 all byte
func wordIndex256(w byte) uint {
	return uint(w)
}

func indexWord256(w uint) byte {
	return byte(w)
}

// wordIndex32to126 空格-~ 0-9 a-z A-Z 符号等
func wordIndex32to126(w byte) uint {
	return uint(w) - ' '
}

func indexWord32to126(w uint) byte {
	return byte(w) + ' '
}
