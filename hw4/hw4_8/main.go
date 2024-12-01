package main

import (
	"fmt"
	"log"
	"strconv"
)

func binconv(num int) string {
	res := ""
	for num > 0 {
		res = strconv.Itoa(num%2) + res
		num /= 2
	}
	return res
}

func setIndexOne(num, i int) int {
	a := 1 << i
	return num | a
}

func setIndexZero(num, i int) int {
	a := ^(1 << i)
	return num & a
}

func invertGivenBit(num, i int) int {
	if (num>>i)%2 == 0 {
		num = setIndexOne(num, i)
	} else {
		num = setIndexZero(num, i)
	}
	return num
}

func main() {
	num, i := 0, 0
	_, err := fmt.Scan(&num, &i)
	if err != nil {
		log.Fatal(err)
	}
	println(binconv(num))

	println(binconv(invertGivenBit(num, i)))
	println(invertGivenBit(num, i))
}
