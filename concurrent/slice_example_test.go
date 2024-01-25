package concurrent

import (
	"log"
	"math"
	"sync"
	"testing"
)

func TestGoSliceAppend(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if math.Mod(float64(i), 2) == 0 {
				AppendSlice(true)
			} else {
				AppendSlice(false)
			}
		}()
	}
	wg.Wait()

	log.Println(len(slice))
}

func TestGoSliceAppendWithLock(t *testing.T) {
	wg := new(sync.WaitGroup)
	mux := new(sync.Mutex)
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if math.Mod(float64(i), 2) == 0 {
				mux.Lock()
				AppendSlice(true)
				mux.Unlock()
			} else {
				mux.Lock()
				AppendSlice(false)
				mux.Unlock()
			}
		}()
	}
	wg.Wait()

	log.Println(len(slice))
}
