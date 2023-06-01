package ArrayAndString

import (
	"fmt"
	"testing"
)

func TestExceptSelf(t *testing.T) {
	self := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Println(self)
}
