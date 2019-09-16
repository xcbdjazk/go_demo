package demo1
import (
	"fmt"
	"testing"
)
// go test -v 
// go test -v cal_test.go cal.go
// go test -v -test.run TestAdd
func TestAdd(t *testing.T)  {
	sum := add(3)
	if sum !=3 {
		t.Fatalf("测试错误 实际值=%v, 期望值=%v", sum, 3)
	}
	fmt.Println("测试完成")
}