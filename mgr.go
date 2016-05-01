package main

type Mgr struct {
	seq   int
	root  *Node
	node0 *Node
	node1 *Node
	ctx   *Lnk
	nn    int
}

func makeMgr() *Mgr {
	mgr := Mgr{seq: -1}
	mgr.root = mgr.makeNode(0)
	mgr.node0 = mgr.makeNode(1)
	mgr.node1 = mgr.makeNode(1)
	mgr.nn = 0
	mgr.link(mgr.root, mgr.node0, mgr.node0)
	mgr.link(mgr.root, mgr.node1, mgr.node1)
	mgr.node0.val = "0"
	mgr.node1.val = "1"

	return &mgr
}

func (mgr *Mgr) readBits(m []byte) {
	cur := mgr.root
	for _, v := range m {
		node := mgr.getBitNode(v)
		cur = mgr.getNextNode(cur, node)
	}
}
