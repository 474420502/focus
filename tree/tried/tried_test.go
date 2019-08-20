package tried

import (
	"bytes"
	"encoding/gob"
	"os"
	"sort"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/Pallinder/go-randomdata"
)

func TestTried_NewWith(t *testing.T) {
	var tried *Tried
	var wordsCollection []string
	var wordsList [][]string
	var triedList []*Tried

	triedList = append(triedList, NewWithWordType(WordIndexLower))
	wordsList = append(wordsList, []string{"adazx", "assdfhgnvb", "ewqyiouyasdfmzvxz"})

	triedList = append(triedList, NewWithWordType(WordIndexUpper))
	wordsList = append(wordsList, []string{"ADFSZ", "DEFASEWRQWER", "GFHJERQWREWTNBVFGFH"})

	triedList = append(triedList, NewWithWordType(WordIndexUpperLower))
	wordsList = append(wordsList, []string{"adazxAZDSAFASZRETHGFTUIPK", "assdfhgDSFGnvb", "yaXZLMPOIQsdGHFfmFBzvxz"})

	triedList = append(triedList, NewWithWordType(WordIndexUpperDigital))
	wordsList = append(wordsList, []string{"AZ3428934470193", "ZPQPDEK09876543629812", "AZEWIRU0192456FDEWR9032"})

	triedList = append(triedList, NewWithWordType(WordIndexLowerDigital))
	wordsList = append(wordsList, []string{"az3428934470193", "zpqwe0987654362sf9812", "az21301az09azdstr540"})

	triedList = append(triedList, NewWithWordType(WordIndexUpperLowerDigital))
	wordsList = append(wordsList, []string{"azAZ09", "aRGFDSFDSzAasdZ06789", "A28374JHFudfsu09qwzzdsw874FDSAZfer"})

	triedList = append(triedList, NewWithWordType(WordIndex256))
	wordsList = append(wordsList, []string{"21`3tcdbxcfhyop8901zc[]\\'/?()#$%^&**! 09-阿萨德发生的官方说的对符合规定", "符号!@$*#))(#*", "╜╝╞╟╠╡╢╣╤╥╦╧╨╩╪╫╬╭╮╯╰╱╲╳▁▂▃▄▅▆▇█ ▉ ▊▋▌▍▎▏"})

	triedList = append(triedList, NewWithWordType(WordIndex32to126))
	wordsList = append(wordsList, []string{" 21`3tcdbxcfhyop8901zc[]\\'/?()#$%^&**!  ", "AZaz09~ dys!@#$)(*^$#", "<>.,?/"})

	for i := 0; i < len(triedList); i++ {
		tried = triedList[i]
		wordsCollection = wordsList[i]
		for _, words := range wordsCollection {
			tried.Put(words)

			if tried.Get(words) == nil {
				t.Error("should be not nil the type is ", tried.wiStore.Type)
			}
		}
		// t.Error(tried.WordsArray())
	}
}

func TestTried_String(t *testing.T) {
	var tried *Tried
	var wordsCollection []string
	var wordsList [][]string
	var triedList []*Tried

	triedList = append(triedList, NewWithWordType(WordIndexLower))
	wordsList = append(wordsList, []string{"adazx", "assdfhgnvb", "ewqyiouyasdfmzvxz"})

	triedList = append(triedList, NewWithWordType(WordIndexUpper))
	wordsList = append(wordsList, []string{"ADFSZ", "DEFASEWRQWER", "GFHJERQWREWTNBVFGFH"})

	triedList = append(triedList, NewWithWordType(WordIndexDigital))
	wordsList = append(wordsList, []string{"093875239457", "09123406534", "0912340846"})

	triedList = append(triedList, NewWithWordType(WordIndexUpperLower))
	wordsList = append(wordsList, []string{"adazxAZDSAFASZRETHGFTUIPK", "assdfhgDSFGnvb", "yaXZLMPOIQsdGHFfmFBzvxz"})

	triedList = append(triedList, NewWithWordType(WordIndexUpperDigital))
	wordsList = append(wordsList, []string{"AZ3428934470193", "ZPQPDEK09876543629812", "AZEWIRU0192456FDEWR9032"})

	triedList = append(triedList, NewWithWordType(WordIndexLowerDigital))
	wordsList = append(wordsList, []string{"az3428934470193", "zpqwe0987654362sf9812", "az21301az09azdstr540"})

	triedList = append(triedList, NewWithWordType(WordIndexUpperLowerDigital))
	wordsList = append(wordsList, []string{"azAZ09", "aRGFDSFDSzAasdZ06789", "A28374JHFudfsu09qwzzdsw874FDSAZfer"})

	triedList = append(triedList, NewWithWordType(WordIndex256))
	wordsList = append(wordsList, []string{"21`3tcdbxcfhyop8901zc[]\\'/?()#$%^&**!\x01 09-213", "!@$*#))(#*", `\/213dsfsdf`})

	triedList = append(triedList, NewWithWordType(WordIndex32to126))
	wordsList = append(wordsList, []string{" 21`3tcdbxcfhyop8901zc[]\\'/?()#$%^&**!  ", "AZaz09~ dys!@#$)(*^$#", "<>.,?/"})

	for i := 0; i < len(triedList); i++ {
		tried = triedList[i]
		wordsCollection = wordsList[i]
		for _, words := range wordsCollection {
			tried.Put(words)
			if tried.Get(words) == nil {
				t.Error("should be not nil the type is ", tried.wiStore.Type)
			}
		}
		sort.Slice(wordsCollection, func(i, j int) bool {
			if wordsCollection[i] < wordsCollection[j] {
				return true
			}
			return false
		})

		resultArray := tried.WordsArray()
		sort.Slice(resultArray, func(i, j int) bool {
			if resultArray[i] < resultArray[j] {
				return true
			}
			return false
		})
		result1 := spew.Sprint(resultArray)
		result2 := spew.Sprint(wordsCollection)
		if result1 != result2 {
			t.Error(result1, " != ", result2)
		}
		// t.Error(tried.WordsArray())
	}
}

func TestTried_PutAndGet1(t *testing.T) {
	tried := New()

	tried.Put(("asdf"))
	tried.Put(("hehe"), "hehe")
	tried.Put(("xixi"), 3)

	var result interface{}

	result = tried.Get("asdf")
	if result != tried {
		t.Error("result should be 3")
	}

	result = tried.Get("xixi")
	if result != 3 {
		t.Error("result should be 3")
	}

	result = tried.Get("hehe")
	if result != "hehe" {
		t.Error("result should be hehe")
	}

	result = tried.Get("haha")
	if result != nil {
		t.Error("result should be nil")
	}

	result = tried.Get("b")
	if result != nil {
		t.Error("result should be nil")
	}
}

func TestTried_Traversal(t *testing.T) {
	tried := New()
	tried.Put("asdf")
	tried.Put(("abdf"), "ab")
	tried.Put(("hehe"), "hehe")
	tried.Put(("xixi"), 3)

	var result []interface{}
	tried.Traversal(func(idx uint, v interface{}) bool {
		// t.Error(idx, v)
		result = append(result, v)
		return true
	})

	if result[0] != "ab" {
		t.Error(result[0])
	}

	if result[1] != tried {
		t.Error(result[1])
	}

	if result[2] != "hehe" {
		t.Error(result[2])
	}

	if result[3] != 3 {
		t.Error(result[3])
	}
}

func TesStoreData(t *testing.T) {
	var l []string
	const N = 1000000
	for i := 0; i < N; i++ {
		var content []rune
		for c := 0; c < randomdata.Number(5, 15); c++ {
			char := randomdata.Number(0, 26) + 'a'
			content = append(content, rune(byte(char)))
		}
		l = append(l, (string(content)))
	}

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(l)
	lbytes := result.Bytes()
	f, _ := os.OpenFile("tried.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	f.Write(lbytes)
}

func Load() []string {
	var result []string
	f, _ := os.Open("tried.log")
	gob.NewDecoder(f).Decode(&result)
	return result
}

func BenchmarkTried_Put(b *testing.B) {

	var data []string
	b.N = 1000000
	count := 10

	// for i := 0; i < b.N; i++ {
	// 	var content []rune
	// 	for c := 0; c < randomdata.Number(5, 15); c++ {
	// 		char := randomdata.Number(0, 26) + 'a'
	// 		content = append(content, rune(byte(char)))
	// 	}
	// 	data = append(data, (string(content)))
	// }

	data = Load()

	b.ResetTimer()
	b.N = b.N * count
	for c := 0; c < count; c++ {
		tried := New()
		for _, v := range data {
			tried.Put(v)
		}
	}
}

func BenchmarkTried_Get(b *testing.B) {
	b.StopTimer()
	var data []string
	b.N = 1000000
	count := 10

	// for i := 0; i < b.N; i++ {
	// 	var content []rune
	// 	for c := 0; c < randomdata.Number(5, 15); c++ {
	// 		char := randomdata.Number(0, 26) + 'a'
	// 		content = append(content, rune(byte(char)))
	// 	}
	// 	data = append(data, string(content))
	// }
	data = Load()

	b.N = b.N * count

	tried := New()
	for _, v := range data {
		tried.Put(v)
	}

	b.StartTimer()
	for c := 0; c < count; c++ {
		for _, v := range data {
			tried.Get(v)
		}
	}
}
