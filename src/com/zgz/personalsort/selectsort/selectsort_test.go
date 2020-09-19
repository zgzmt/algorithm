package selectsort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSelectsort01(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	type test struct {
		input []int
		want  []int
	}
	test1 := []test{}
	for i := 0; i < 5; i++ {
		t := test{}
		z := rand.Intn(150)
		for z >= 0 {
			t.input = append(t.input, rand.Intn(300))
			z--
		}
		test1 = append(test1, t)
		//fmt.Println(t)
	}
	for _, tw := range test1 {
		tw.want = make([]int, len(tw.input))
		copy(tw.want, tw.input)
		//	append(tw.input)
		sort.Ints(tw.want)
		got := Selectsort01(tw.input)
		if !reflect.DeepEqual(got, tw.want) {
			t.Errorf("excepted:%v,got:%v", tw.want, got)
		}
	}

}
