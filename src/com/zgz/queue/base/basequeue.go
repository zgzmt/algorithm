package base

func init() {
	pqueue := new(Queue)
	pqueue.info = -1
	pqueue.next = nil
	queue = pqueue
	Head = pqueue
}

type Queue struct {
	next *Queue
	info int
}

var Head *Queue
var queue *Queue

func InQueue(a int) {
	p := new(Queue)
	p.info = a
	p.next = nil
	queue.next = p
	queue = queue.next
}

func OutQueue() (ret int) {
	q := Head
	ret = Head.next.info
	Head = Head.next
	q.next = nil
	return
}

func Run(arr []int) (ret []int) {
	for i := 0; i < len(arr); i++ {
		InQueue(arr[i])
	}
	//j := 0
	for Head.next != nil {
		ret = append(ret, OutQueue())
	}
	return
}
