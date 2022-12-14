package timer

import (
	"fmt"
	"time"
)

func TimeIt(cnt int, f func()) {
	times := make([]int64, 0, cnt)
	for i := 0; i < cnt; i++ {
		start := time.Now()
		f()
		times = append(times, time.Since(start).Microseconds())
	}

	var avg int64 = 0
	for _, t := range times {
		avg += t
	}

	fmt.Println("avg: ", float64(avg)/float64(len(times)))
}
