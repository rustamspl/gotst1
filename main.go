package main

func main() {
	mgr := makeMgr()
	//mgr.readBits([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})4

	mgr.readStr("а").print()
	mgr.readStr("б").print()
	mgr.readStr("аб").print()
	mgr.readStr("wабабабd").print()
	mgr.readStr("а").print()

	//mgr.readBits([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
