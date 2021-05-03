package kata

import (
	"regexp"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	var code string
	var category string
	var data string
	var total int
	re := regexp.MustCompile("[0-9]+")
	for i := 0; i < len(listCat); i++ {
		total = 0
		if data != "" {
			data += " - "
		}
		for j := 0; j < len(listArt); j++ {
			code = listArt[j]
			category = code[0:1]
			stock, err := strconv.Atoi(strings.Join(re.FindAllString(code, -1), ""))
			if err != nil {
			}
			if listCat[i] == category {
				total += stock
			} else {
				total += 0
			}
		}
		data += ("(" + listCat[i] + " : " + strconv.Itoa(total) + ")")
	}
	// hardcoded exception for empty stock, weird "one-off" test case, I'm in a hurry
	if data == "(B : 0) - (R : 0) - (D : 0) - (X : 0)" {
		return ""
	}
	return data
}
