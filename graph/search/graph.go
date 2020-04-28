package gsearch

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/474420502/focus/tree/heap"
)

// Graph 图
type Graph struct {
	// direction   []*Graph
	// flag        int
	isDebug bool

	weightFunc func(npather *Pather, graph *Graph) int
	weightHeap *heap.Tree

	steps      int
	stepslimit int

	patherMap [][]*Pather // 记录最小的路径pather 比这个大而且来过的都可以cut
	tile      [][]*Point

	srart *Point
	end   *Point

	dimX  int
	dimY  int
	tsize int // dimX * dimY
	bsize int // bit size

	blockflag float64
}

// Point 点
type Point struct {
	X, Y int
	Flag byte
}

func weightCompare(x1, x2 interface{}) int {
	p1, p2 := x1.(*Pather), x2.(*Pather)
	if p1.weight > p2.weight {
		return 1
	}
	return -1
}

// New2D 一个graph. 必须指定维度数据
func New2D(dx, dy int) *Graph {
	g := &Graph{}
	g.setDimension(dx, dy)
	g.tsize = g.dimX * g.dimY
	g.bsize = (g.tsize + 1) / 8
	g.weightHeap = heap.New(weightCompare)
	g.weightFunc = SimpleWeight

	g.tile = make([][]*Point, dy)
	for y := 0; y < dy; y++ {
		g.tile[y] = make([]*Point, dx)
		for x := 0; x < dx; x++ {
			point := &Point{X: x, Y: y}
			g.tile[y][x] = point
		}
	}

	g.patherMap = make([][]*Pather, dy)
	for y := 0; y < dy; y++ {
		g.patherMap[y] = make([]*Pather, dx)
	}

	g.stepslimit = g.tsize * 128
	return g
}

// New2DFromBlockFile 一个graph. 必须指定BlockFile
func New2DFromBlockFile(path string) *Graph {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var sdatalist [][]string
	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		sdata := regexp.MustCompile(`\d+`).FindAllString(string(line), -1)
		if sdata != nil {
			sdatalist = append(sdatalist, sdata)
		}
	}
	dy := len(sdatalist) // 维度
	dx := len(sdatalist[0])
	graph := New2D(dx, dy)

	for y, sdata := range sdatalist {
		tiltx := graph.tile[y]
		for x, s := range sdata {
			blockvalue, err := strconv.ParseInt(s, 16, 8)
			if err != nil {
				panic(err)
			}
			tiltx[x].Flag = byte(blockvalue)
		}
	}

	// graph.CountBlocksFlag()
	return graph
}

// CountBlocksFlag 计算blocks的比例, 便于做估价处理
// func (graph *Graph) CountBlocksFlag() {
// 	blocks := 0
// 	zeroSize := graph.tsize
// 	for y := 0; y < graph.dimY; y++ {
// 		for x := 0; x < graph.dimX; x++ {
// 			msize := y*graph.dimY + x
// 			v := graph.tile[msize]
// 			if v > 0 {
// 				blocks++
// 			} else {
// 				zeroCount := 0
// 				for _, p := range []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

// 					cx := x + p.x
// 					cy := y + p.y
// 					if cx < 0 || cy < 0 {
// 						break
// 					}

// 					if cx >= graph.dimX || cy >= graph.dimY {
// 						break
// 					}

// 					cmsize := cy*graph.dimY + cx
// 					if graph.tile[cmsize] == 0 {
// 						zeroCount++
// 					}
// 				}

// 				if zeroCount == 4 {
// 					zeroSize -= 2
// 				}
// 			}
// 		}
// 	}

// 	graph.blockflag = float64(blocks) / float64(zeroSize)
// }

// setDimension 初始化维度
func (graph *Graph) setDimension(dx, dy int) {
	graph.dimX = dx
	graph.dimY = dy
}

// SetTimeoutSteps 设置起点 结束点
func (graph *Graph) SetTimeoutSteps(steps int) {
	graph.stepslimit = steps
}

// SetBlock 设置障碍　完后需要调用　CountBlocksFlag() 让估价更加准确.
func (graph *Graph) SetBlock(x, y int, v byte) {
	graph.tile[y][x].Flag = v
}

// SetBlockFromFile 设置障碍从文件中 内部已经调用 CountBlocksFlag
func (graph *Graph) SetBlockFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var sdatalist [][]string
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		sdata := regexp.MustCompile(`\d+`).FindAllString(string(line), -1)
		if sdata != nil {
			sdatalist = append(sdatalist, sdata)
		}
	}
	dy := len(sdatalist) // 维度
	dx := len(sdatalist[0])

	if !(dy == graph.dimY && dx == graph.dimX) {
		panic("file dim X, Y is not equal graph tile")
	}

	for y, sdata := range sdatalist {
		tiltx := graph.tile[y]
		for x, s := range sdata {
			blockvalue, err := strconv.ParseInt(s, 16, 8)
			if err != nil {
				panic(err)
			}
			tiltx[x].Flag = byte(blockvalue)
		}
	}
	// graph.CountBlocksFlag()
}

// SetTarget 设置起点 结束点
func (graph *Graph) SetTarget(sx, sy, ex, ey int) {

	graph.srart = graph.tile[sy][sx]
	graph.end = graph.tile[ey][ex]

	// graph.end.msize = graph.end.y*graph.dimX + graph.end.x

	gdata := NewBitmap2D(graph.dimX, graph.dimY)
	paths := make([]*Point, 0)

	pather := newPather(graph.srart, gdata, paths, 0)
	graph.weightHeap.Put(pather)

	graph.cutByMinPath(pather)
}

// GetStep 执行的步数
func (graph *Graph) GetStep() int {
	return graph.steps
}

// Search 执行搜索
func (graph *Graph) Search() (*Pather, bool) {

	defer func() {
		if graph.isDebug {
			log.Println("blockflag", graph.blockflag, "search heap size:", graph.weightHeap.Size())
		}
		graph.weightHeap.Clear()
	}()

	graph.steps = 0
	for !graph.weightHeap.Empty() {
		ipather, _ := graph.weightHeap.Pop()
		pather := ipather.(*Pather)
		if graph.Traversing(pather) {
			return pather, true
		}

		graph.steps++
		if graph.steps >= graph.stepslimit {
			log.Println("超时找不到路径", graph.steps, graph.weightHeap.Size())
			return pather, false
		}
	}
	return nil, false
}

// Traversing 遍历结果
func (graph *Graph) Traversing(pather *Pather) bool {
	if pather.pos == graph.end {
		if graph.isDebug {
			graph.debugShow(pather)
		}
		// log.Println("finish")
		return true
	}

	pather.path = append(pather.path, pather.pos)
	pather.bits.SetBit(pather.pos.X, pather.pos.Y, 1)

	// graph.myDebug(pather)

	graph.left(pather)
	graph.right(pather)
	graph.up(pather)
	graph.down(pather)

	return false
}

// SetWeight 设置估价函数
func (graph *Graph) SetWeight(weight func(npather *Pather, graph *Graph) int) {
	graph.weightFunc = weight
}

// Clear 设置估价函数
func (graph *Graph) Clear() {
	graph.steps = 0
	graph.blockflag = 0

	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			graph.patherMap[y][x] = nil
		}
	}
	graph.weightHeap.Clear()
}

func (graph *Graph) evaluate(npather *Pather, pather *Pather) {
	if graph.cut(npather, pather.bits) {
		return
	}

	npather.bits = CopyFrom(pather.bits)

	npather.path = make([]*Point, len(pather.path))
	copy(npather.path, pather.path)

	npather.weight = graph.weightFunc(npather, graph)
	graph.weightHeap.Put(npather)
}

func (graph *Graph) cutByMinPath(pather *Pather) bool {
	checkpoint := pather.pos
	if checkpather := graph.patherMap[checkpoint.Y][checkpoint.X]; checkpather != nil {
		if checkpather.GetPathLen() > pather.GetPathLen() {
			graph.patherMap[checkpoint.Y][checkpoint.X] = pather
		} else {
			return true
		}
	} else {
		graph.patherMap[checkpoint.Y][checkpoint.X] = pather
	}
	return false
}

func (graph *Graph) cutUnessential(npather *Pather, bits *Bitmap2D) bool {
	count := 0
	for _, p := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

		x := npather.pos.X + p[0] // p[0] == `x`
		y := npather.pos.Y + p[1] // p[0] == `y`
		if x < 0 || y < 0 {
			continue
		}

		if x >= graph.dimX || y >= graph.dimY {
			continue
		}

		if bits.GetBit(x, y) > 0 {
			count++
			if count > 1 {
				return true
			}
		}
	}
	return false
}

func (graph *Graph) cut(npather *Pather, bits *Bitmap2D) bool {

	pinfo := npather.pos.Flag
	if pinfo&0b00000001 > 0 { // 障碍物
		return true
	}

	if bits.GetBit(npather.pos.X, npather.pos.Y) > 0 {
		return true
	}

	if graph.cutByMinPath(npather) {
		return true
	}

	if graph.cutUnessential(npather, bits) {
		return true
	}

	return false
}

func (graph *Graph) left(pather *Pather) {
	leftx := pather.pos.X - 1
	if leftx < 0 {
		return
	}

	npather := &Pather{pos: graph.tile[pather.pos.Y][leftx]}
	graph.evaluate(npather, pather)
}

func (graph *Graph) right(pather *Pather) {
	rightx := pather.pos.X + 1
	if rightx >= graph.dimX {
		return
	}

	npather := &Pather{pos: graph.tile[pather.pos.Y][rightx]}
	graph.evaluate(npather, pather)
}

func (graph *Graph) up(pather *Pather) {
	upy := pather.pos.Y - 1
	if upy < 0 {
		return
	}

	npather := &Pather{pos: graph.tile[upy][pather.pos.X]}
	graph.evaluate(npather, pather)
}

func (graph *Graph) down(pather *Pather) {
	downy := pather.pos.Y + 1
	if downy >= graph.dimY {
		return
	}

	npather := &Pather{pos: graph.tile[downy][pather.pos.X]}
	graph.evaluate(npather, pather)
}

func (graph *Graph) debugShow(pather *Pather) {
	pather.bits.SetBit(pather.pos.X, pather.pos.Y, 1)

	content := "\n"
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			content += fmt.Sprintf("%1d ", pather.bits.GetBit(x, y))
		}
		content += "\n"
	}
	log.Println(content)

	content = "\n"
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			content += fmt.Sprintf("%02x ", graph.tile[y][x].Flag)
		}
		content += "\n"
	}
	log.Println(content)
}

func (graph *Graph) myDebug(pather *Pather) {

	plen := len(pather.path)
	if plen >= int(math.Sqrt(float64(graph.dimY*graph.dimY+graph.dimX*graph.dimX))+1) && graph.steps >= 504288 {

		content := "\n"
		for y := 0; y < graph.dimY; y++ {
			for x := 0; x < graph.dimX; x++ {
				content += fmt.Sprintf("%1d ", pather.bits.GetBit(x, y))
			}
			content += "\n"
		}
		log.Println(content)
	}
}
