package main

import (
	"strconv"
	"strings"
)

func main() {
	subtractProductAndSum(214)
}

func subtractProductAndSum(n int) int {
	str := strconv.Itoa(n)
	arr := strings.Split(str, "")
	product := 1
	sum := 0
	for _, i := range arr {
		num, _ := strconv.Atoi(i)
		product *= num
		sum += num
	}
	return product - sum
}
