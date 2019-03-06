package examples

import (
	"fmt"
	"sync/atomic"
	"time"
)

func AtomicCounters() {
	fmt.Println("**Examples Atomic Counters**")
	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops", opsFinal)

}
