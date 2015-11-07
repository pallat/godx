package main

import (
	"flag"
	"fmt"
	"go-x2go"
	"io/ioutil"
)

var file string
var output string

func init() {
	flag.StringVar(&file, "f", "example.xml", "-f=example.xml")
	flag.StringVar(&output, "o", "example.go", "-o=example.go")
}

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	x2go := x2go.New(b)

	ioutil.WriteFile(output, []byte(x2go.String()), 0666)
}
