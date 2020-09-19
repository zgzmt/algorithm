package selectsort

import "fmt"

func Selectsort01(arr []int) []int {
	i := len(arr)
	var temp int
	for j := 0; j < i-1; j++ {
		temp = arr[j]
		for z := j + 1; z < i; z++ {
			if temp > arr[z] {
				temp = arr[z]
				arr[z] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
	return arr
}
