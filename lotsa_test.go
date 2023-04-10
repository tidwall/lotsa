package lotsa

import (
	"math/rand"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

func TestOps(t *testing.T) {
	Output = os.Stdout
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

func TestTime(t *testing.T) {
	Output = os.Stdout
	var threads = 4
	var expectedDuration = 1 * time.Millisecond
	var total int64

	startTs := time.Now()
	Time(expectedDuration, threads, func(threadRand *rand.Rand, thread int) {
		atomic.AddInt64(&total, 1)
	})
	actualDuration := time.Since(startTs)

	if expectedDuration.Round(time.Millisecond) != actualDuration.Round(time.Millisecond) {
		t.Fatal("invalid duration")
	}
}
