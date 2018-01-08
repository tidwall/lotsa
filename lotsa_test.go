package lotsa

import (
	"sync/atomic"
	"testing"
)

func TestOps(t *testing.T) {
	var threads = 4
	var N = threads * 10000
	var total int64
	Ops(N, threads, func(i, thread int) {
		atomic.AddInt64(&total, 1)
	})
	if total != int64(N) {
		t.Fatal("invalid total")
	}
}
