package main

import "fmt"

type Node struct {
	size, id int //todo: del
	plus     map[*Node]*Node
	minus    map[*Node]*Node
}

func (n *Node) bits() []byte {

	if n.size > 1 {
		for x := range n.minus {
			if x.size > 0 {
				return append(n.minus[x].bits(), x.bits()...)
			}
		}
	}
	if n.id == 0 {
		return []byte{0}
	} else {
		return []byte{1}
	}
}
func (n *Node) str() string {
	m := n.bits()
	r := ""
	var (
		i uint
		j uint
		l uint
		t rune
	)
	l = uint(len(m))
	for i = 0; i < l; i += 16 {
		t = 0
		for j = 0; j < 16; j++ {
			t = t | (rune(m[i+j]) << j)
		}
		r = r + string(t)
	}
	return r
}
func (n *Node) print() {
	r := ""
	if n.size > 16 {
		for x := range n.minus {
			if x.size > 0 {
				r = n.minus[x].str() + "+" + x.str()
				break
			}
		}
	}
	fmt.Println(n.id, n.str(), ":", r)
}
