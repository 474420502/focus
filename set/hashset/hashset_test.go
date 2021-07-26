package hashset

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"testing"
)

func loadTestData() []int {
	log.SetFlags(log.Lshortfile)

	data, err := ioutil.ReadFile("../../l.log")
	if err != nil {
		log.Println(err)
	}
	var l []int
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&l)
	return l
}

func TestAdd(t *testing.T) {

	set := New()
	for i := 0; i < 10; i++ {
		set.Add(i)
	}

	if set.Size() != 10 {
		t.Error("size is not equals to 10")
	}
}

func TestRemove(t *testing.T) {
	set := New()

	for i := 0; i < 10; i++ {
		set.Add(i)
	}

	for i := 0; i < 9; i++ {
		set.Remove(i)
	}

	if set.Size() != 1 {
		t.Error("size is not equals to 0")
	}

	if set.Values()[0] != 9 {
		t.Error("remain is not 9")
	}
}
