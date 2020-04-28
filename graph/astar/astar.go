package astar

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

	weight     func(nparam *Param, graph *Graph) int
	weightHeap *heap.Tree

	steps      int
	stepslimit int

	paramMap [][]*Param // 记录最小的路径Param 比这个大而且来过的都可以cut
	tile     [][]*Point

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
	p1, p2 := x1.(*Param), x2.(*Param)
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
	g.weight = SimpleWeight

	g.tile = make([][]*Point, dy)
	for y := 0; y < dy; y++ {
		g.tile[y] = make([]*Point, dx)
		for x := 0; x < dx; x++ {
			point := &Point{X: x, Y: y}
			g.tile[y][x] = point
		}
	}

	g.paramMap = make([][]*Param, dy)
	for y := 0; y < dy; y++ {
		g.paramMap[y] = make([]*Param, dx)
	}

	g.stepslimit = (dx*dx + dy*dy) * 512
	return g
}

// New2DFromBlockFile 一个graph. 必须指定BlockFile 内部已经调用 CountBlocksFlag
func New2DFromBlockFile(path string) *Graph {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var sdatalist [][]string
	reader := bufio.NewReader(f)
	dy := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		sdata := regexp.MustCompile(`\d+`).FindAllString(line, -1)
		if sdata != nil {
			dy++
			sdatalist = append(sdatalist, sdata)
		}
	}
	dy++ // 维度
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
	dy := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		sdata := regexp.MustCompile(`\d+`).FindAllString(line, -1)
		if sdata != nil {
			dy++
			sdatalist = append(sdatalist, sdata)
		}
	}
	dy++ // 维度
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

	param := newParam(graph.srart, gdata, paths, 0)
	graph.weightHeap.Put(param)

	graph.cutAllByMinPath(param)
}

// GetStep 执行的步数
func (graph *Graph) GetStep() int {
	return graph.steps
}

// Search 执行搜索
func (graph *Graph) Search() (*Param, bool) {

	defer func() {
		if graph.isDebug {
			log.Println("blockflag", graph.blockflag, "search heap size:", graph.weightHeap.Size())
		}
		graph.weightHeap.Clear()
	}()

	graph.steps = 0
	for !graph.weightHeap.Empty() {
		iparam, _ := graph.weightHeap.Pop()
		param := iparam.(*Param)
		if graph.Traversing(param) {
			return param, true
		}

		graph.steps++
		if graph.steps >= graph.stepslimit {
			log.Println("超时找不到路径", graph.steps, graph.weightHeap.Size())
			return param, false
		}
	}
	return nil, false
}

// Traversing 遍历结果
func (graph *Graph) Traversing(param *Param) bool {
	if param.pos == graph.end {
		if graph.isDebug {
			graph.debugShow(param)
		}
		// log.Println("finish")
		return true
	}

	param.paths = append(param.paths, param.pos)
	param.bits.SetBit(param.pos.X, param.pos.Y, 1)

	// graph.myDebug(param)

	graph.left(param)
	graph.right(param)
	graph.up(param)
	graph.down(param)

	return false
}

// SetWeight 设置估价函数
func (graph *Graph) SetWeight(weight func(nparam *Param, graph *Graph) int) {
	graph.weight = weight
}

func (graph *Graph) evaluate(nparam *Param, param *Param) {
	nparam.bits = CopyFrom(param.bits)

	nparam.paths = make([]*Point, len(param.paths))
	copy(nparam.paths, param.paths)

	nparam.weight = graph.weight(nparam, graph)
	graph.weightHeap.Put(nparam)
}

func (graph *Graph) cutAllByMinPath(param *Param) bool {
	checkpoint := param.pos
	if checkParam := graph.paramMap[checkpoint.Y][checkpoint.X]; checkParam != nil {
		if len(checkParam.paths) > len(param.paths) {
			graph.paramMap[checkpoint.Y][checkpoint.X] = param
		} else {
			return true
		}
	} else {
		graph.paramMap[checkpoint.Y][checkpoint.X] = param
	}
	return false
}

func (graph *Graph) cutUnessential(nparam *Param, bits *Bitmap2D) bool {
	count := 0
	for _, p := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

		x := nparam.pos.X + p[0] // p[0] == `x`
		y := nparam.pos.Y + p[1] // p[0] == `y`
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

func (graph *Graph) cut(nparam *Param, bits *Bitmap2D) bool {

	if graph.cutAllByMinPath(nparam) {
		return true
	}

	if graph.cutUnessential(nparam, bits) {
		return true
	}

	return false
}

func (graph *Graph) left(param *Param) {
	leftx := param.pos.X - 1
	if leftx < 0 {
		return
	}

	nparam := &Param{pos: graph.tile[param.pos.Y][leftx]}
	if param.bits.GetBit(nparam.pos.X, nparam.pos.Y) > 0 {
		return
	}

	if graph.cut(nparam, param.bits) {
		return
	}

	pinfo := nparam.pos.Flag
	if pinfo&0b00000001 > 0 { // 障碍物
		return
	}

	graph.evaluate(nparam, param)
}

func (graph *Graph) right(param *Param) {
	rightx := param.pos.X + 1
	if rightx >= graph.dimX {
		return
	}

	nparam := &Param{pos: graph.tile[param.pos.Y][rightx]}
	if param.bits.GetBit(nparam.pos.X, nparam.pos.Y) > 0 {
		return
	}

	if graph.cut(nparam, param.bits) {
		return
	}

	pinfo := nparam.pos.Flag
	if pinfo&0b00000001 > 0 { // 障碍物
		return
	}

	graph.evaluate(nparam, param)
}

func (graph *Graph) up(param *Param) {
	upy := param.pos.Y - 1
	if upy < 0 {
		return
	}

	nparam := &Param{pos: graph.tile[upy][param.pos.X]}
	if param.bits.GetBit(nparam.pos.X, nparam.pos.Y) > 0 {
		return
	}

	if graph.cut(nparam, param.bits) {
		return
	}

	pinfo := nparam.pos.Flag
	if pinfo&0b00000001 > 0 { // 障碍物
		return
	}

	graph.evaluate(nparam, param)
}

func (graph *Graph) down(param *Param) {
	downy := param.pos.Y + 1
	if downy >= graph.dimY {
		return
	}

	nparam := &Param{pos: graph.tile[downy][param.pos.X]}
	if param.bits.GetBit(nparam.pos.X, nparam.pos.Y) > 0 {
		return
	}

	if graph.cut(nparam, param.bits) {
		return
	}

	pinfo := nparam.pos.Flag
	if pinfo&0b00000001 > 0 { // 障碍物
		return
	}

	graph.evaluate(nparam, param)
}

func (graph *Graph) debugShow(param *Param) {
	param.bits.SetBit(param.pos.X, param.pos.Y, 1)

	content := "\n"
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			content += fmt.Sprintf("%1d ", param.bits.GetBit(x, y))
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

func (graph *Graph) myDebug(param *Param) {

	plen := len(param.paths)
	if plen >= int(math.Sqrt(float64(graph.dimY*graph.dimY+graph.dimX*graph.dimX))+1) && graph.steps >= 504288 {

		content := "\n"
		for y := 0; y < graph.dimY; y++ {
			for x := 0; x < graph.dimX; x++ {
				content += fmt.Sprintf("%1d ", param.bits.GetBit(x, y))
			}
			content += "\n"
		}
		log.Println(content)
	}
}
