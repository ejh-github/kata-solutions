package kata

import "fmt"

func BouncingBall(h, bounce, window float64) int {
	fmt.Println(h, bounce, window)
	if !(h > 0 && (bounce > 0 && bounce < 1) && window < h) {
		return -1
	} else if h < window {
		return 1
	} else {
		return 2 + BouncingBall((h*bounce), bounce, window)
	}
}
