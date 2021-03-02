package kata

import(
	"strconv"
	"strings"
)

func CountBits(number uint) int {
	binary := strconv.FormatInt(int64(number), 2)
	count := strings.Count(string(binary), "1")
	return count
}
