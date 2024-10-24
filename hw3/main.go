package main

import (
	"fmt"
	"github.com/gwkeo/Tower_Back_HW/hw3/cli"
	"github.com/gwkeo/Tower_Back_HW/hw3/uniq"
	"log"
)

func main() {
	content, attributes, err := cli.GetAttributes()
	if err != nil {
		log.Fatal(err)
	}
	result, ok := uniq.Uniq(content, attributes)
	if ok != nil {
		log.Fatal(ok)
	}
	fmt.Println(result)
}
