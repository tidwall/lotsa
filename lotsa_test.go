package lotsa

import (
	"bytes"
	"strings"
	"sync/atomic"
	"testing"
)

func TestOps(t *testing.T) {
	var threads = 4
	var N = threads * 10000
	var total int64
	var wr bytes.Buffer
	Ops(N, threads, &wr, func(i, thread int) {
		atomic.AddInt64(&total, 1)
	})
	if total != int64(N) {
		t.Fatal("invalid total")
	}
	if !strings.Contains(wr.String(), commaize(N)) {
		t.Fatal("bad writer")
	}
}
