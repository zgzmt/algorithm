package languangeinsite

import "fmt"

type bitTree struct {
	val        int
	lef, right *bitTree
}

//递归先序便利二叉树
func bitTreePreOrderTra(root *bitTree) {
	if root == nil {
		return
	}
	fmt.Printf("%d", root.val)
	bitTreePreOrderTra(root.lef)
	bitTreePreOrderTra(root.right)
}

//递归中序遍历二叉树
func bitTreeMidOrderTra(root *bitTree) {
	if root == nil {
		return
	}
	bitTreeMidOrderTra(root.lef)
	fmt.Printf("%d", root.val)
	bitTreeMidOrderTra(root.right)
}

//递归后序遍历二叉树
func bitTreePosOrderTra(root *bitTree) {
	if root == nil {
		return
	}
	bitTreePosOrderTra(root.lef)
	bitTreePosOrderTra(root.right)
	fmt.Printf("%d", root.val)
}
