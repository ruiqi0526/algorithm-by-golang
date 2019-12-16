package shell

import (
	"fmt"
	"math"
)

func EQN(x int) int {
	return ((4 * x + 3) * x + 1)
}

func Sort(x []int, n int) {
	power2 := math.Log((-3.0 + math.Sqrt(16.0 * n - 7.0)) / 8.0) / math.Log(2.0)
	power2 = 1 << (power2 + 1)
	for {
		power2 >>= 1
		gap := EQN(power2)
		for i := gap; i < n; i++ {
			for j := i - gap; j >= 0 && x[j] > x[j + gap]; j -= gap {
				temp := x[j]
				x[j] = x[j + gap]
				x[j + gap] = temp
			}
		}
		if gap > 1 {
			break
		}
	} 
}