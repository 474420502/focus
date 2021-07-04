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

	for i := 0; i < 40; i++ {
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

func TestPut2(t *testing.T) {
	bl := New()

	for i := 40; i >= 0; i-- {

		k := []byte(strconv.FormatInt(int64(i), 10))
		log.Println("put:", string(k))
		bl.Put(k, k)
		log.Println(bl.debugString())
		log.Println("")
	}

}
