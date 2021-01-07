package languangeinsite

//切片模仿栈
func slinceStack(){
	stack := make([]int,0)//创建栈
	//进栈
	stack = append(stack,1)
	//出栈
	var _ = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	//判断栈空
	if len(stack) ==0{}
}

//切片模仿队列
