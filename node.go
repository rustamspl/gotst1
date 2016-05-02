package main

//	"fmt"

type Node struct {
	size, id int //todo: del
	plus     map[*Node]*Node
	minus    map[*Node]*Node
}
