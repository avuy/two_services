package main

import "testing"

func benchmarkForSender(n int, b *testing.B) {
	client(n)
	//b.StopTimer()
}

func Benchmark_Check_1_Site(b *testing.B)        { benchmarkForSender(1, b) }
func Benchmark_Check_10_Sites(b *testing.B)      { benchmarkForSender(10, b) }
func Benchmark_Check_100_Sites(b *testing.B)     { benchmarkForSender(100, b) }
func Benchmark_Check_1000_Sites(b *testing.B)    { benchmarkForSender(1000, b) }
func Benchmark_Check_10000_Sites(b *testing.B)   { benchmarkForSender(10000, b) }
func Benchmark_Check_100000_Sites(b *testing.B)  { benchmarkForSender(100000, b) }
//func Benchmark_Check_1000000_Sites(b *testing.B) { benchmarkForSender(1000000, b) }
