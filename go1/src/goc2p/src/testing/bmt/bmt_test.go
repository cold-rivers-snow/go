package bmt

import (
	"testing"
	"time"
)

func Benchmark(b *testing.B) {
	customTimerTag := false
	if customTimerTag {
		b.StopTimer()
	}
	b.SetBytes(12345678)		//单次操作被处理的字节的数量。
	time.Sleep(time.Second)
	if customTimerTag {
		b.StartTimer()
	}
}
