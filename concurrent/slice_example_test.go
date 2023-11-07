package concurrent

import (
	"math"
	"sync"
	"testing"
)

func TestGoSlice(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if math.Mod(float64(i), 2) == 0 {
				SetSlice()
			} else {
				GetSlice()
			}
		}()
	}
	wg.Wait()
}
