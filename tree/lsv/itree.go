package lsv

import (
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// INode 用于索引的节点
type INode struct {
	family [3]*INode
	size   int
	tree   *DTree
}

// NewINode 生成inode节点
func NewINode() *INode {
	inode := &INode{}
	inode.size = 1
	inode.tree = newDataTree(compareRunes)
	return inode
}

// ITree 用于索引的树
type ITree struct {
	root    *INode
	limit   int
	Compare func(s1, s2 []rune) int
}

// New 生成一颗索引树
func New(Compare func(s1, s2 []rune) int) *ITree {
	return &ITree{Compare: Compare, limit: 100}
}

// Put return bool
func (tree *ITree) Put(key, value []rune) (isInsert bool) {

	// node := &INode{key: key, value: value, size: 1}
	if tree.root == nil {
		tree.root = NewINode()
		return tree.root.tree.Put(key, value)
	}

	for cur := tree.root; ; {

		if cur.size > 8 {
			factor := cur.size >> 3 // or factor = 1
			ls, rs := cur.family[1].size, cur.family[2].size
			if rs >= (ls<<1)+factor || ls >= (rs<<1)+factor {
				tree.ifixSize(cur, ls, rs)
			}
		}

		c := tree.Compare(key, cur.tree.feature.key)
		switch {
		case c < 0:
			if cur.family[1] == nil {

				if cur.tree.root.size >= tree.limit {
					// 子树的节点分解　操作
					lspilt := cur.tree.root.family[1]
					rspilt := cur.tree.root.family[2]
					rspilt.family[0] = nil //清空右节点的父类, rsplit为根
					lspilt.family[0] = nil // 上

					tempRoot := cur.tree.root
					tempRoot.size = 1
					for i := 1; i < len(tempRoot.family); i++ {
						tempRoot.family[i] = nil
					}

					var icur *INode
					ilnode := NewINode()
					ilnode.tree.root = lspilt
					ilnode.tree.putfeature(tempRoot)

					cur.family[1] = ilnode
					ilnode.family[0] = cur
					icur = cur

					cur.tree.root = rspilt //主根替换 右树

					for temp := icur; temp != nil; temp = temp.family[0] {
						temp.size++
					} // 往上加+1 达到每个节点能统计size正确

					// 调整3节点失衡的情况
					if icur.family[0] != nil && icur.family[0].size == 3 {
						if icur.family[0].family[1] == nil {
							tree.ilrrotate3(icur.family[0])
						} else {
							tree.irrotate3(icur.family[0])
						}
					}
				}

				return cur.tree.Put(key, value)
			}
			cur = cur.family[1]
		case c > 0:
			if cur.family[2] == nil {

				if cur.tree.root.size >= tree.limit {
					lspilt := cur.tree.root.family[1]
					rspilt := cur.tree.root.family[2]
					rspilt.family[0] = nil //清空右节点的父类, rsplit为根
					lspilt.family[0] = nil // 上

					tempRoot := cur.tree.root
					tempRoot.size = 1
					for i := 1; i < len(tempRoot.family); i++ {
						tempRoot.family[i] = nil
					}

					// cur.family
					var icur *INode
					irnode := NewINode()
					irnode.tree.root = rspilt
					irnode.tree.feature = cur.tree.feature

					cur.family[2] = irnode
					irnode.family[0] = cur
					icur = cur

					for temp := icur; temp != nil; temp = temp.family[0] {
						temp.size++
					} // 往上加+1 达到每个节点能统计size正确

					// 调整3节点失衡的情况
					if icur.family[0] != nil && icur.family[0].size == 3 {
						if icur.family[0].family[2] == nil {
							tree.irlrotate3(icur.family[0])
						} else {
							tree.ilrotate3(icur.family[0])
						}
					}

					cur.tree.root = lspilt
					cur.tree.putfeature(tempRoot)
				}

				return cur.tree.Put(key, value)
			}
			cur = cur.family[2]
		default:
			// c == 0 而且满足插入限制, 分离出来, 分左右节点
			return cur.tree.Put(key, value)
		}

	}
}

func (tree *ITree) String() {

}

func (tree *ITree) ifixSize(cur *INode, ls, rs int) {
	if ls > rs {
		llsize, lrsize := igetChildrenSize(cur.family[1])
		if lrsize > llsize {
			tree.irlrotate(cur)
		} else {
			tree.irrotate(cur)
		}
	} else {
		rlsize, rrsize := igetChildrenSize(cur.family[2])
		if rlsize > rrsize {
			tree.ilrrotate(cur)
		} else {
			tree.ilrotate(cur)
		}
	}
}

func (tree *ITree) ilrrotate3(cur *INode) {
	const l = 2
	const r = 1

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	cur.family[r] = mov
	mov.family[0] = cur

	cur.family[l] = movparent
	movparent.family[r] = nil

	cur.family[r] = mov
	mov.family[0] = cur

	cur.family[l].size = 1
}

func (tree *ITree) ilrrotate(cur *INode) {

	const l = 2
	const r = 1

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	if mov.family[l] != nil {
		movparent.family[r] = mov.family[l]
		movparent.family[r].family[0] = movparent
		//movparent.family[r].child = l
	} else {
		movparent.family[r] = nil
	}

	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
		//mov.family[l].child = l
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	cur.family[r] = mov
	mov.family[0] = cur

	movparent.size = igetChildrenSumSize(movparent) + 1
	mov.size = igetChildrenSumSize(mov) + 1
	cur.size = igetChildrenSumSize(cur) + 1
}

func (tree *ITree) irlrotate3(cur *INode) {
	const l = 1
	const r = 2

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	cur.family[r] = mov
	mov.family[0] = cur

	cur.family[l] = movparent
	movparent.family[r] = nil

	cur.family[r] = mov
	mov.family[0] = cur

	// cur.size = 3
	// cur.family[r].size = 1
	cur.family[l].size = 1
}

func (tree *ITree) irlrotate(cur *INode) {

	const l = 1
	const r = 2

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	if mov.family[l] != nil {
		movparent.family[r] = mov.family[l]
		movparent.family[r].family[0] = movparent
	} else {
		movparent.family[r] = nil
	}

	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	cur.family[r] = mov
	mov.family[0] = cur

	movparent.size = igetChildrenSumSize(movparent) + 1
	mov.size = igetChildrenSumSize(mov) + 1
	cur.size = igetChildrenSumSize(cur) + 1
}

func (tree *ITree) irrotate3(cur *INode) {
	const l = 1
	const r = 2
	// 1 right 0 left
	mov := cur.family[l]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	cur.family[r] = mov

	cur.family[l] = mov.family[l]
	cur.family[l].family[0] = cur

	mov.family[l] = nil

	mov.size = 1
}

func (tree *ITree) irrotate(cur *INode) {

	const l = 1
	const r = 2
	// 1 right 0 left
	mov := cur.family[l]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	//  mov.family[l]不可能为nil
	mov.family[l].family[0] = cur

	cur.family[l] = mov.family[l]

	// 解决mov节点孩子转移的问题
	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.family[r] = mov

	mov.size = igetChildrenSumSize(mov) + 1
	cur.size = igetChildrenSumSize(cur) + 1
}

func (tree *ITree) ilrotate3(cur *INode) {
	const l = 2
	const r = 1
	// 1 right 0 left
	mov := cur.family[l]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	cur.family[r] = mov

	cur.family[l] = mov.family[l]
	cur.family[l].family[0] = cur

	mov.family[l] = nil

	mov.size = 1
}

func (tree *ITree) ilrotate(cur *INode) {

	const l = 2
	const r = 1
	// 1 right 0 left
	mov := cur.family[l]

	mov.tree, cur.tree = cur.tree, mov.tree //交换值达到, 相对位移

	//  mov.family[l]不可能为nil
	mov.family[l].family[0] = cur

	cur.family[l] = mov.family[l]

	// 解决mov节点孩子转移的问题
	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.family[r] = mov

	mov.size = igetChildrenSumSize(mov) + 1
	cur.size = igetChildrenSumSize(cur) + 1
}

func igetChildrenSumSize(cur *INode) int {
	return igetSize(cur.family[1]) + igetSize(cur.family[2])
}

func igetChildrenSize(cur *INode) (int, int) {
	return igetSize(cur.family[1]), igetSize(cur.family[2])
}

func igetSize(cur *INode) int {
	if cur == nil {
		return 0
	}
	return cur.size
}

func (tree *ITree) ifixSizeWithRemove(cur *INode) {
	for cur != nil {
		cur.size--
		if cur.size > 8 {
			factor := cur.size >> 3 // or factor = 1
			ls, rs := igetChildrenSize(cur)
			if rs >= (ls<<1)+factor || ls >= (rs<<1)+factor {
				tree.ifixSize(cur, ls, rs)
			}
		} else if cur.size == 3 {
			if cur.family[1] == nil {
				if cur.family[2].family[1] == nil {
					tree.ilrotate3(cur)
				} else {
					tree.ilrrotate3(cur)
				}
			} else if cur.family[2] == nil {
				if cur.family[1].family[2] == nil {
					tree.irrotate3(cur)
				} else {
					tree.irlrotate3(cur)
				}
			}
		}
		cur = cur.family[0]
	}
}

func ioutputfordebug(node *INode, prefix string, isTail bool, str *string) {

	if node.family[2] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		ioutputfordebug(node.family[2], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	suffix := "("
	parentv := ""
	if node.family[0] == nil {
		parentv = "nil"
	} else {
		parentv = spew.Sprint(string(node.family[0].tree.root.key[0:3]), strconv.Itoa(node.family[0].tree.root.size))
	}
	suffix += parentv + "|" + spew.Sprint(node.size) + ")"
	*str += spew.Sprint(string(node.tree.root.key[0:3]), strconv.Itoa(node.tree.root.size)) + suffix + "\n"

	if node.family[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		ioutputfordebug(node.family[1], newPrefix, true, str)
	}
}

func (tree *ITree) debugString() string {
	str := "LSV\n"
	if tree.root == nil {
		return str + "nil"
	}
	ioutputfordebug(tree.root, "", true, &str)
	return str
}
