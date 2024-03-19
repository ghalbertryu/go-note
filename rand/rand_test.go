package main

import (
	cRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestCryptoRand(t *testing.T) {
	p := 1.0 / 6.0
	for i := 0; i < 10; i++ {
		counter := make(map[int64]int)
		num := 1000000
		for i := 0; i < num; i++ {
			n, err := cRand.Int(cRand.Reader, big.NewInt(6))
			if err != nil {
				panic(err)
			}

			numInt := n.Int64()
			if count, ok := counter[numInt]; !ok {
				counter[numInt] = 1
			} else {
				counter[numInt] = count + 1
			}
		}

		proportion := make(map[int64]float64, len(counter))
		for k, count := range counter {
			proportion[k] = float64(count) / float64(num)
		}

		proportionGreater := make(map[int64]bool, len(counter))
		for k, prop := range proportion {
			proportionGreater[k] = prop > p
		}
		fmt.Println("count", counter)
		fmt.Println("prob", proportion)
		fmt.Println("greater", proportionGreater)
		fmt.Println()
	}
}

func TestRand(t *testing.T) {
	counter := make(map[int]int)
	num := 1000000
	for i := 0; i < num; i++ {
		numInt := rand.Intn(6)
		if count, ok := counter[numInt]; !ok {
			counter[numInt] = 1
		} else {
			counter[numInt] = count + 1
		}
	}

	proportion := make(map[int]float64, len(counter))
	for k, count := range counter {
		proportion[k] = float64(count) / float64(num)
	}
	fmt.Println(counter)
	fmt.Println(proportion)

}

func TestCR(t *testing.T) {
	num := 50
	for j := 0; j < num; j++ {
		n, _ := cRand.Int(cRand.Reader, big.NewInt(6))
		fmt.Print(n, " ")
	}
}

func TestR(t *testing.T) {
	num := 50
	for j := 0; j < num; j++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		v := r.Intn(6)
		fmt.Print(v, " ")
	}
}

func TestDiceMap(t *testing.T) {
	n, err := cRand.Int(cRand.Reader, big.NewInt(6))
	if err != nil {
		panic(err)
	}
	diceNum := n.Int64() + 1
	print(diceNum)
}
