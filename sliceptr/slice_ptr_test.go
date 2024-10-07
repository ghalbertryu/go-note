package sliceptr

import (
	"fmt"
	"log"
	"testing"
	"time"
	"unsafe"
)

func TestSlicePtr(t *testing.T) {
	// 创建一个切片
	s := []int{1, 2, 3}

	// 打印切片本身的地址
	fmt.Printf("Slice struct address: %p\n", &s)

	// 打印切片中指向的底层数组地址（即切片的第一个元素地址）
	fmt.Printf("Slice underlying array pointer: %p\n", s)

	// 打印切片中每个元素的地址
	for i := 0; i < len(s); i++ {
		fmt.Printf("Address of element %d: %p\n", i, &s[i])
	}

	// 打印底层数组的起始地址（等效于切片第一个元素的地址）
	fmt.Printf("Underlying array address (unsafe.Pointer): %p\n", unsafe.Pointer(&s[0]))
}

func TestSliceSetup(t *testing.T) {
	s := buildIntSlice()
	fmt.Printf("Slice struct address: %p\n", &s)
	fmt.Printf("Slice underlying array pointer: %p\n", s)

	fmt.Println("TestSliceSetup 1", s)
	time.Sleep(1500)
	fmt.Println("TestSliceSetup 2", s)

	time.Sleep(3 * time.Second)
}

func TestByteSlice(t *testing.T) {
	bytes := buildByteSlice()
	fmt.Printf("Slice struct address bytes: %p\n", &bytes)
	fmt.Printf("Slice underlying array pointer bytes: %p\n", bytes)

	str := string(bytes)
	fmt.Printf("Slice struct address str: %p\n", &str)
	//fmt.Printf("Slice underlying array pointer str: %p\n", str)
	log.Println(str)

	bytes2 := []byte(str)
	fmt.Printf("Slice struct address bytes2: %p\n", &bytes2)
	fmt.Printf("Slice underlying array pointer bytes2: %p\n", bytes2)
	log.Println(bytes2)
}
