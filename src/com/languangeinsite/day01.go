package languangeinsite

import (
	"fmt"
	"sort"
	"strconv"
)

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
func slinceQueue(){
	//创建队列
	queue := make([]int,0)
	//进队
	queue = append(queue,10)
	//出队
	var _ = queue[0]
	queue = queue[1:]
	//队列判空
	if len(queue) == 0 {}
}

//字典的用法
func useMap(){
	//创建map
	m:= make(map[string]int)
	//赋值
	m["hello"] = 1
	//判断某个值是否存以及获取值
	if v,ok := m["hello"]; ok{
		fmt.Println(v)
	}else{
		fmt.Println("不存在该值")
	}
	//删除map中的某个键值
	delete(m,"hello")
	//遍历
	for k,v := range m{
		fmt.Printf("key=%s,value=%d\n",k,v)
	}
}
//标准库排序
func standSort(){
	//整型排序
	intarry := []int{1,5,7,45,24,68,21,8,16,65}
	sort.Ints(intarry)
	//字符串排序
	stringarry := []string{"adb","dec","acd"}
	sort.Strings(stringarry)
	//自定义排序
	sort.Slice(intarry,func(i,j int)bool{
		return intarry[i]>intarry[j]
	})
}
//复制
func useCopy(){
	//删除切片的地i个值；可以用copy将第i+1复制钱前面
	i := 5
	a := make([]int,10)
	copy(a[i:],a[i+1:])
	//make指定长度，则可以索引赋值
	a1 := make([]int,i)
	a1[1]=1
	//make长度为0则需要使用append来赋值
	a2 := make([]int,0)
	a2 = append(a2,3)
}
//常用技巧
func useuse(){
	//字符串转数字
	str1 := "123456"
	num,_ := strconv.Atoi(str1)
	fmt.Printf("type:%V,value:%d",num,num)

	//整形转字符
	str := strconv.Itoa(123)
	fmt.Printf(str)

}
