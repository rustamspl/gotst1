package main

import (
	"fmt"
)

type Node struct {
	size     int
	id       int    //todo: del
	val      string //todo: del
	plus     map[*Node]*Node
	minus    map[*Node]*Node
	maxMinus *Node
	minMinus *Node
}

func (mgr *Mgr) getBitNode(bit byte) *Node {
	if bit == 0 {
		return mgr.node0
	} else {
		return mgr.node1
	}
}

func (mgr *Mgr) makeNode(size int) *Node {
	node := Node{size: size}
	node.id = mgr.seq
	mgr.seq++
	node.plus = make(map[*Node]*Node)
	node.minus = make(map[*Node]*Node)
	return &node
}

func (mgr *Mgr) link(a *Node, b *Node, c *Node) {

	a.plus[b] = c
	c.minus[b] = a

	if b.size <= (c.size>>1) && ((c.maxMinus == nil) || (c.maxMinus.size < b.size)) { //
		c.maxMinus = b
	}

	if b.size > 0 && ((c.minMinus == nil) || (c.minMinus.size > b.size)) { //
		c.minMinus = b
	}
	fmt.Println(c.id, "=", a.id, "+", b.id, "; ", c.size, "=", a.size, "+", b.size)
}

func (mgr *Mgr) getNextNode(a *Node, b *Node) *Node {
	mgr.nn += 1
	//	if mgr.nn > 1000 {
	//		panic("aaa")
	//	}
	if c, has := a.plus[b]; has {
		return c
	} else {
		if a.maxMinus != nil && (a.size-b.size) > 1 {
			//fmt.Println(a.id, " > ", b.id)
			subNode := mgr.getNextNode(a.maxMinus, b)
			return mgr.getNextNode(a.minus[a.maxMinus], subNode)

		}
		if b.minMinus != nil && (b.size-a.size) > 1 {
			//fmt.Println(a.id, " < ", b.id)
			subNode := mgr.getNextNode(a, b.minus[b.minMinus])
			return mgr.getNextNode(subNode, b.minMinus)

		}
		newNode := mgr.makeNode(a.size + b.size)
		mgr.link(a, b, newNode)
		return newNode

	}
}
