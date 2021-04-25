package main

func transpose(A [][]int) [][]int {
	var (
		b [][]int
		lenA int = len(A)
		lenB int = len(A[0])
	)

	for i:=0; i<lenB; i++{
		var a []int
		for j:=0; j<lenA; j++{
			a = append(a, A[i][j])
		}
		b = append(b,a)
	}
	return b
}

func main() {
	var a = [][]int{
		[]int{1,2,3},
		[]int{1,2,3},
	}
	transpose(a)
}