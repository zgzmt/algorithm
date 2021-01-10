package main

import (
	//"com.languangeinsite"
	//"com/languangeinsite"
	//"fmt"
	"fmt"
)
type bitTreeNode struct {
	val int
	lef ,rright *bitTreeNode
}
var root *bitTreeNode
func init(){
	var p *bitTreeNode
	p = &bitTreeNode{
		val:    0,
		lef:    nil,
		rright: nil,
	}
	had := p
	for i := 1 ; i< 5 ; i++ {
		p.lef = &bitTreeNode{
			val:    i,
			lef:    nil,
			rright: nil,
		}
		p.rright = &bitTreeNode{
			val:    i+10,
			lef:    nil,
			rright: nil,
		}
		p = p.lef
	}
	fmt.Printf("head类型：%T\n",had)
	fmt.Printf("headd地址：%p\n",had)
	fmt.Printf("栈内p类型：%T\n",p)
	fmt.Printf("栈内p指针地址：%p\n",p)
	root = had
}

func main() {
	//fmt.Println("Hello world")
	//languangeinsite.StrStr("Hello world","Hello")
	fmt.Printf("root的类型:%T\n",root)
	fmt.Printf("root的类型:%p\n",root)
	var testarr []int
	if testarr == nil {
		fmt.Println("yes")
	}

}
