package main

type Queue struct {
	array   [4]int
	maxSize int // 大小
	front   int // 队首
	rear    int // 队尾
}

func (q Queue) add(i int) (err error) {

	return

}
func (q Queue) push() (err error) {

	return
}

func (q Queue) size() int {
	return q.front - q.rear
}

func main() {

}
