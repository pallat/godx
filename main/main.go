package main

import (
	"flag"
	"fmt"
	"go-x2go"
	"io/ioutil"
)

var file string

func init() {
	flag.StringVar(&file, "file", "example.xml", "-file=example.xml")
}

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	x2go := x2go.New(b)
	fmt.Println(x2go.String())
}
