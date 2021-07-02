package blist

import (
	"log"
	"strconv"
	"testing"

	"github.com/Pallinder/go-randomdata"
)

func TestPut(t *testing.T) {
	bl := New()

	var dict = make(map[int]bool)
	for i := 0; i < 50; i++ {
		// for _, n := range []int{42, 15, 35, 34} {
		n := randomdata.Number(0, 100)
		if _, ok := dict[n]; !ok {
			dict[n] = true
		} else {
			i--
			continue
		}

		k := []byte(strconv.FormatInt(int64(n), 10))
		log.Println("put:", string(k))
		bl.Put(k, k)
		log.Println(bl.debugString())
		log.Println("")
	}

}
