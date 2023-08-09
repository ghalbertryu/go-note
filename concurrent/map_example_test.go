package concurrent

import (
	"math"
	"sync"
	"testing"
)

func TestGoMap(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if math.Mod(float64(i), 2) == 0 {
				Set()
			} else {
				Get()
			}
		}()
	}
	wg.Wait()
}

func TestGoSyncMap(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if math.Mod(float64(i), 2) == 0 {
				SetSyncMap()
			} else {
				GetSyncMap()
			}
		}()
	}
	wg.Wait()
}
