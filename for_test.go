package focus

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"testing"

	randomdata "github.com/Pallinder/go-randomdata"
)

const CompartorSize = 2000000
const NumberMax = 500000000

func TestSave(t *testing.T) {

	f, err := os.OpenFile("l.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println(userBytes)

	var l []int

	// for i := 0; len(l) < 1000; i++ {
	// 	v := randomdata.Number(0, 65535)
	// 	l = append(l, v)
	// }

	//m := make(map[int]int)
	for i := 0; len(l) < CompartorSize; i++ {
		v := randomdata.Number(0, NumberMax)
		// if _, ok := m[v]; !ok {
		// 	m[v] = v
		l = append(l, v)
		// }
	}

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(l)
	lbytes := result.Bytes()
	f.Write(lbytes)

}

func LoadTestData() []int {
	data, err := ioutil.ReadFile("../l.log")
	if err != nil {
		log.Println(err)
	}
	var l []int
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&l)
	return l
}
