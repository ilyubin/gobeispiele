package convert

import (
	"fmt"
	"strconv"
	"testing"
)

//BenchmarkString-4      	50000000	        29.6 ns/op
//BenchmarkFormatInt-4   	50000000	        27.3 ns/op
//BenchmarkItoa-4        	50000000	        29.0 ns/op
//BenchmarkSprint-4      	20000000	        83.0 ns/op

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(int32(i))
	}
}

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(int64(i), 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(int(i))
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(i)
	}
}
