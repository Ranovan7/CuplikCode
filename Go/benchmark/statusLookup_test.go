package main

import "testing"

var M map[int]string = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Elevel",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
}

// Switch statement
func getValueSwitch(key int) string {
	switch key {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"
	case 15:
		return "Fifteen"
	default:
		return "Unknown"
	}
}

// Hashmap (Map)
func getValueMap(key int) string {
	if value, ok := M[key]; ok {
		return value
	}
	return "Unknown"
}

// Benchmark for switch
func BenchmarkStatusLookupSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getValueSwitch(i % 16)
	}
}

// Benchmark for map
func BenchmarkStatusLookupMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getValueMap(i % 16)
	}
}
