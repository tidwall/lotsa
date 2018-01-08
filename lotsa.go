package lotsa

import (
	"fmt"
	"io"
	"sync"
	"time"
)

// Ops executed a number of operations over a multiple goroutines.
// count is the number of operations.
// threads is the number goroutines.
// output is a writer that will be used to print the results.
// op is the operation function
func Ops(count, threads int, output io.Writer, op func(i, thread int)) {
	var start time.Time
	if output != nil {
		start = time.Now()
	}
	var wg sync.WaitGroup
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		s, e := count/threads*i, count/threads*(i+1)
		if e > count {
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
		tp := "s"
		if threads == 1 {
			tp = ""
		}
		fmt.Fprintf(output, "%s ops over %d thread%s in %s (%s ops/sec)\n",
			commaize(count), threads, tp, dur, commaize(int(float64(count)/dur.Seconds())))
	}
}
func commaize(n int) string {
	nstr1, nstr := fmt.Sprintf("%d", n), ""
	for i, j := len(nstr1)-1, 0; i >= 0; i, j = i-1, j+1 {
		nstr = string(nstr1[i]) + nstr
		if i != 0 && j%3 == 2 {
			nstr = "," + nstr
		}
	}
	return nstr
}
