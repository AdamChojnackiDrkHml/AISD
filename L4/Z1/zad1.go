package main

type tree interface {
	insert(value string)
	load(uri string) bool
	delete(value string) bool
	find(value string) bool
	min() (int bool)
	max() (int bool)
	successor(value string) (int bool)
	inorder()
}

type BSTnode struct {
	Value    string
	LeftSon  *BSTnode
	RightSon *BSTnode
	Parent   *BSTnode
}

func (BSTnode) insert(Val string) {
	newNode := BSTnode{Value: Val}

}

func main() {

}
