package _range

import (
	"log"
	"testing"
	"time"
)

func TestRangeIssue(t *testing.T) {
	collet := buildCollection()
	log.Println(collet)

	for _, dto := range collet {
		go func() {
			log.Println(dto)
		}()
	}

	time.Sleep(3 * time.Second)
}

func TestSolve1(t *testing.T) {
	collet := buildCollection()
	log.Println(collet)

	for _, dto := range collet {
		go func(obj any) {
			log.Println(obj)
		}(dto)
	}

	time.Sleep(3 * time.Second)
}

func TestSolve2(t *testing.T) {
	collet := buildCollection()
	log.Println(collet)

	for i, _ := range collet {
		dto := collet[i]
		go func() {
			log.Println(dto)
		}()
	}

	time.Sleep(3 * time.Second)
}
