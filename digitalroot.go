package kata

import (
	"fmt"
	"strconv"
)

func DigitalRoot(n int) int {
	var num int
	var err error
	sum := 0
	fmt.Println(n)
	nstring := strconv.Itoa(n)
	for i := 0; i < len(nstring); i++ {
		num, err = strconv.Atoi(nstring[i : i+1])
		fmt.Println(num)
		sum += num
		if err != nil {
		}
	}
	fmt.Println(sum)
	if sum > 9 {
		return DigitalRoot(sum)
	} else {
		return sum
	}
}

/* This code does the same thing as above, much cleaner and more beautiful, optimal solution.
package kata

func DigitalRoot(n int) int {
    return (n - 1) % 9 + 1
}

*/
