package squaredup

import (
	"log"
	"testing"
)

func TestCase1(t *testing.T) {
	su := New(3, 3)
	vlen := 1
	su.start = NewGird(3, 3, vlen)
	su.start.SetGirdByString(3, vlen,
		`
		1 2 3
		4 0 5
		7 8 6
		`,
	)
	su.randomGird(su.start, 100)
	log.Println(su.start.GetGirdString(su.dimY, su.dimX, su.vlen))

	su.end = NewGird(3, 3, vlen)
	su.end.SetGirdByString(3, vlen,
		`
		1 2 3
		4 5 6
		7 8 0
		`,
	)
	su.Search()
}

func TestCase2(t *testing.T) {

	dim := 4
	vlen := 1

	su := New(dim, dim)
	su.start = NewGird(dim, dim, vlen)
	su.start.SetGirdByString(dim, vlen,
		`
		1  2  3  4
		5  6  11 7
		9 14 10 12
		0 13 8 15
		`,
	)
	su.randomGird(su.start, 200)
	log.Println(su.start.GetGirdString(su.dimY, su.dimX, su.vlen))
	su.end = NewGird(dim, dim, vlen)
	su.end.SetGirdByString(dim, vlen,
		`
		1  2  3  4
		5  6  7  8
		9 10 11 12
		13 14 15 0
		`,
	)
	su.Search()
}

func TestCase3(t *testing.T) {

	dim := 5
	vlen := 1

	su := New(dim, dim)
	su.start = NewGird(dim, dim, vlen)
	su.start.SetGirdByString(dim, vlen,
		`
		1  2  3 4 5
		6  7  8 9 10 
		11 12 13 14 15
		16 17 18 19 0
		21 22 23 24 20
		`,
	)
	su.randomGird(su.start, 200)
	log.Println(su.start.GetGirdString(su.dimY, su.dimX, su.vlen))
	su.end = NewGird(dim, dim, vlen)
	su.end.SetGirdByString(dim, vlen,
		`
		1  2  3 4 5
		6  7  8 9 10 
		11 12 13 14 15
		16 17 18 19 20
		21 22 23 24 0
		`,
	)
	su.Search()
}
