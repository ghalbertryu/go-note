package sliceptr

import (
	"fmt"
	"time"
)

func buildIntSlice() []int {
	s := []int{1, 2, 3}
	fmt.Printf("Slice struct address: %p\n", &s)
	fmt.Printf("Slice underlying array pointer: %p\n", s)

	// 變更 slice 內元素資料
	time.Sleep(1000)
	go func(s []int) {
		s[0] = 8888
		fmt.Println("buildIntSlice", s)
	}(s)
	return s
}
