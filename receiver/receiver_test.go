package main

import "testing"

func benchmarkForReceiver(n int, b *testing.B) {
	b.StopTimer()
}

func Benchmark_Check_1_Site(b *testing.B)        { benchmarkForReceiver(1, b) }
func Benchmark_Check_10_Sites(b *testing.B)      { benchmarkForReceiver(10, b) }
func Benchmark_Check_100_Sites(b *testing.B)     { benchmarkForReceiver(100, b) }
func Benchmark_Check_1000_Sites(b *testing.B)    { benchmarkForReceiver(1000, b) }
func Benchmark_Check_10000_Sites(b *testing.B)   { benchmarkForReceiver(10000, b) }
func Benchmark_Check_100000_Sites(b *testing.B)  { benchmarkForReceiver(100000, b) }
func Benchmark_Check_1000000_Sites(b *testing.B) { benchmarkForReceiver(1000000, b) }
