package lotsa

import "sync"

// Ops executed a number of operations over a multiple goroutines.
// count is the number of operations.
// threads is the number goroutines.
// op is the operation function
func Ops(count, threads int, op func(i, thread int)) {
	var wg sync.WaitGroup
	wg.Add(threads)
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
}
