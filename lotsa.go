package lotsa

import (
	"fmt"
	"io"
	"sync"
	"time"
)

// Output is used to print elased time and ops/sec
var Output io.Writer

// Ops executed a number of operations over a multiple goroutines.
// count is the number of operations.
// threads is the number goroutines.
// op is the operation function
func Ops(count, threads int, op func(i, thread int)) {
	var start time.Time
	var wg sync.WaitGroup
	wg.Add(threads)
	output := Output
	if output != nil {
		start = time.Now()
	}
	for i := 0; i < threads; i++ {
		s, e := count/threads*i, count/threads*(i+1)
		if i == threads-1 {
			e = count
		}
		go func(i, s, e int) {
			for j := s; j < e; j++ {
				op(j, i)
			}
			wg.Done()
		}(i, s, e)
	}
	wg.Wait()
	if output != nil {
		dur := time.Since(start)
		fmt.Fprintf(output, "%d ops over %d threads in %.0fms %.0f/sec\n",
			count, threads, dur.Seconds()*1000, float64(count)/dur.Seconds())

	}
}
