package main

import (
	"fmt"
	"io/ioutil"
)

func cat(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", contents)
}

func main() {
	cat("tina.txt")
}
