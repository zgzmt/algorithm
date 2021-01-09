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
//非递归先序遍历二叉树
func preOrderTra(root *bitTree){
	if root ==nil {
		return
	}
	p := root
	var stack []*bitTree
	for p!=nil || stack != nil {
		for p != nil {
			fmt.Print("%d",p.val)
			stack = append(stack,p)
			p = p.lef
		}
			p = stack[len(stack)-1].right
			stack = stack[:len(stack)-1]

	}
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

//中序遍历二叉树，非递归
func midOrderTra(root *bitTree)[]int{
	if root != nil {
		return nil
	}
	p := root
	reslut := make([]int,0)
	stack := make([]*bitTree,0)
	for p !=nil || len(stack)!=0 {
		for p!=nil {
			stack = append(stack,p)
			p = p.lef
		}
			p = stack[len(stack)-1]
			reslut = append(reslut,p.val)
			stack = stack[:len(stack)-1]
			p = p.right

	}
	return reslut
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

//非递归后续遍历二叉树
func posordertra(root *bitTree)[]int {
	if root == nil {
		return nil
	}
	result := make([]int,0)
	stack := make([]*bitTree,0)
	p := root
	var lastNode *bitTree
	if p!=nil || len(stack)!=0{
		for p != nil {
			stack = append(stack,p)
			p = p.lef
		}
		node := stack[len(stack)-1]
		if node.right == nil || node == lastNode {
			stack = stack[:len(stack)-1]
			result = append(result,node.val)
			lastNode = node
		}else{
			p = node.right
		}
	}
	return result
}
