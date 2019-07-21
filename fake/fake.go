package fake

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Phone() string {
	return fmt.Sprintf("799%d", RandomDigits(8))
}

func Email() string {
	return fmt.Sprintf("qaapi.ozon.ru+%d%d@gmail.com", RandomDigits(8), rangeIn(10, 99))
}

func RandomDigits(length int) int64 {
	ms := time.Now().UnixNano() / int64(time.Millisecond)
	a := ((ms - (ms % 10)) / 10) % int64(math.Pow10(length))
	if a < int64(math.Pow10(length-1)) {
		a = a + 5*int64(math.Pow10(length-1))
	}
	return a
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
