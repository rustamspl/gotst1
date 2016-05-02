package main

type Lst struct {
	root *Lnk
	len  int
}

type Lnk struct {
	node       *Node
	next, prev *Lnk
}

func makeLst() *Lst {
	root := &Lnk{}
	root.next = root
	root.prev = root
	lst := Lst{len: 0, root: root}
	return &lst
}

func (lst *Lst) insert(e, at *Lnk) {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	lst.len++
}
func (lst *Lst) remove(e *Lnk) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	lst.len--
}

func (lst *Lst) PushFront(node *Node) {
	lnk := &Lnk{}
	lnk.node = node
	lst.insert(lnk, lst.root)
}
