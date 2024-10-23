package main

import (
	"fmt"
	"github.com/gwkeo/Tower_Back_HW/hw3/uniq"
)

func main() {
	var str []string
	str = append(str, "cake")
	str = append(str, "cake")
	str = append(str, "bake")
	str = append(str, "")
	str = append(str, "cake")
	attributes := &uniq.Attributes{
		ReturnOnlyUniqueLines: true,
		NumberOfCharsToSkip:   1,
	}
	fmt.Println(uniq.Uniq(str, attributes))
}
