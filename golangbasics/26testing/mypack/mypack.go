package mypack

import "sync/atomic"

var count uint64

// Countby2 byt2
func Countby2() uint64 {
	atomic.AddUint64(&count, 2)
	return count
}
