package main

import "fmt"

type Mgr struct {
	seq          int
	node0, node1 *Node
	ctx          *Lst
}

func (mgr *Mgr) makeNode(size int) *Node {
	node := Node{size: size}
	node.id = mgr.seq
	node.plus = make(map[*Node]*Node)
	node.minus = make(map[*Node]*Node)
	mgr.seq++
	return &node
}
func makeMgr() *Mgr {
	mgr := Mgr{seq: 0}
	mgr.node0 = mgr.makeNode(1)
	mgr.node1 = mgr.makeNode(1)
	mgr.ctx = makeLst()
	return &mgr
}
func (mgr *Mgr) getBitNode(bit byte) *Node {
	if bit == 0 {
		return mgr.node0
	} else {
		return mgr.node1
	}
}
func (mgr *Mgr) readBits(m []byte) {
	for _, v := range m {
		mgr.readBit(v)
	}

}
func (mgr *Mgr) readBit(bit byte) {
	node := mgr.getBitNode(bit)
	mgr.ctx.PushFront(node)
	cur := mgr.ctx.root.next
	for i := mgr.ctx.len - 1; i > 0; i-- {
		next := cur.next
		curNode := mgr.linkSameSizeNodes(next.node, cur.node)
		if curNode != nil {
			next.node = curNode
			mgr.ctx.remove(cur)
			cur = next
		} else {
			break
		}
	}

}
func (mgr *Mgr) linkSameSizeNodes(a, b *Node) *Node {
	if c, has := a.plus[b]; has {
		return c
	}

	if a.size == b.size {
		return mgr.makeLink(a, b)
	}

	return nil
}
func (mgr *Mgr) makeLink(a, b *Node) *Node {
	c := mgr.makeNode(a.size + b.size)
	a.plus[b] = c
	c.minus[b] = a
	//fmt.Println(c.id, "=", a.id, "+", b.id)
	return c
}

func (mgr *Mgr) linkCtxNodes() *Node {
	if mgr.ctx.len > 1 {
		cur := mgr.ctx.root.next
		for mgr.ctx.len > 1 {
			next := cur.next

			if c, has := next.node.plus[cur.node]; has {
				next.node = c
			} else {
				next.node = mgr.makeLink(next.node, cur.node)
			}
			mgr.ctx.remove(cur)
			cur = next
		}
	}
	return mgr.ctx.root.next.node

}
func (mgr *Mgr) readStr(s string) *Node {
	mgr.ctx.clear()
	for _, v := range s {
		t := v
		for i := 0; i < 16; i++ {
			bit := byte(t & 1)
			mgr.readBit(bit)
			t = t >> 1
		}
	}
	return mgr.linkCtxNodes()

}
func (mgr *Mgr) printCtx() {
	for i, cur := mgr.ctx.len, mgr.ctx.root.prev; i > 0; i, cur = i-1, cur.prev {
		fmt.Println(cur.node.size, cur.node.id)
	}
}
