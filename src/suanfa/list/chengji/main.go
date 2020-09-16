package main

import "fmt"

func run(l [5]int) {
	var (
		maxVal   int
		minVal   int
		minIndex int
		maxIndex int
	)
	for i, v := range l {
		fmt.Println(i, v)
		if maxVal < v {
			maxVal = v
			maxIndex = i
		}
		if minVal == 0 {
			minVal = v
		}
		if minVal > v {
			minVal = v
			minIndex = i
		}
	}
	inx1 := maxIndex
	inx2 := minIndex
	if maxIndex < minIndex {
		inx1 = minIndex
		inx2 = maxIndex
	}
	fmt.Println(l)
	for i := inx1; i < len(l)-1; i++ {
		l[i] = l[i+1]
	}
	fmt.Println(l)
	for i := inx2; i < len(l)-1; i++ {
		l[i] = l[i+1]
	}

	sumscore := 0
	for i := 0; i < len(l)-2; i++ {
		sumscore += l[i]
	}
	fmt.Println(l)
	fmt.Println("index", maxIndex, maxVal, minIndex, minVal, sumscore)
}

func main() {
	run([5]int{11, 2, 3, 4, 5})
}
