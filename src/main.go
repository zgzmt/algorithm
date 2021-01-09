package main

import (
	//"com.languangeinsite"
	//"com/languangeinsite"
	//"fmt"
	"fmt"
)
type bitTree struct {
	val int
	lef ,rright *bitTree
}
var root * bitTree
func init(){
	var p *bitTree
	p = &bitTree{
		val:    0,
		lef:    nil,
		rright: nil,
	}
	for i := 1 ; i< 5 ; i++ {
		p.lef = &bitTree{
			val:    i,
			lef:    nil,
			rright: nil,
		}
		p.rright = &bitTree{
			val:    i+10,
			lef:    nil,
			rright: nil,
		}
		p = p.lef
	}
	root = p
}

func main() {
	//fmt.Println("Hello world")
	//languangeinsite.StrStr("Hello world","Hello")
	fmt.Printf("root的类型%T\n",root)
}
