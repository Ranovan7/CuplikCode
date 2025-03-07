package main

import (
	"fmt"
	"testing"
)

// Struct with a large payload (to highlight performance differences)
type LargeStruct struct {
	varA1 string
	varB1 uint64
	varC1 float64
	varA2 string
	varB2 uint64
	varC2 float64
	varA3 string
	varB3 uint64
	varC3 float64
}

var ls LargeStruct = LargeStruct{
	varA1: "asfe2efdc23r2ec 23r23ef 2evc ef 2er23cr2f 23r 24f 3erf f23f2c3efc23f 23f2vr23f 23f 23f 22 3f23r2v3rc23r2v3t2 g2v23t2r 23r 23",
	varB1: 1_111_111_111,
	varC1: 0.341231141231,
	varA2: "asfe2efdc23r2ec 23r23ef 2evc ef 2er23cr2f 23r 24f 3erf f23f2c3efc23f 23f2vr23f 23f 23f 22 3f23r2v3rc23r2v3t2 g2v23t2r 23r 23",
	varB2: 1_111_111_111,
	varC2: 0.341231141231,
	varA3: "asfe2efdc23r2ec 23r23ef 2evc ef 2er23cr2f 23r 24f 3erf f23f2c3efc23f 23f2vr23f 23f 23f 22 3f23r2v3rc23r2v3t2 g2v23t2r 23r 23",
	varB3: 1_111_111_111,
	varC3: 0.341231141231,
}

// 🔹 Value Receiver: Returns a copy of the struct
func (ls LargeStruct) ValueMethod() {
	_ = fmt.Sprintf("%s_%d_%f", ls.varA1, ls.varB1, ls.varC1)
	_ = fmt.Sprintf("%s_%d_%f", ls.varA2, ls.varB2, ls.varC2)
	_ = fmt.Sprintf("%s_%d_%f", ls.varA3, ls.varB3, ls.varC3)
}

// 🔹 Pointer Receiver: Modifies the original struct
func (ls *LargeStruct) PointerMethod() {
	_ = fmt.Sprintf("%s_%d_%f", ls.varA1, ls.varB1, ls.varC1)
	_ = fmt.Sprintf("%s_%d_%f", ls.varA2, ls.varB2, ls.varC2)
	_ = fmt.Sprintf("%s_%d_%f", ls.varA3, ls.varB3, ls.varC3)
}

// Benchmark for Value Receiver (Copy)
func BenchmarkPVVValueReceiver(b *testing.B) {
	obj := LargeStruct{} // Create an instance
	for i := 0; i < b.N; i++ {
		obj.ValueMethod()
	}
}

// Benchmark for Pointer Receiver (No Copy)
func BenchmarkPVVPointerReceiver(b *testing.B) {
	obj := LargeStruct{} // Create an instance
	for i := 0; i < b.N; i++ {
		obj.PointerMethod()
	}
}
