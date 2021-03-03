package kata

import "fmt"

func BouncingBall(h, bounce, window float64) int {
	if !(h > 0 && (bounce > 0 && bounce < 1) && window < h) {
		return -1
	} else {
		return 2 + BouncingBall((h*bounce), bounce, window)
	}
}
