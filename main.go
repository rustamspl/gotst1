package main

func main() {
	mgr := makeMgr()
	//mgr.readBits([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})4

	mgr.readStr("привет")
	mgr.printCtx()
	mgr.readStr("привет")
	mgr.printCtx()
	//mgr.readBits([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
