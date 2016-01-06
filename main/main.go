package main

import (
	"flag"
	"fmt"
	"godx"
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

	x := godx.New(b)
	ioutil.WriteFile(output, []byte(x.String()), 0666)

	// b, err = format.Source([]byte(x.String()))

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// ioutil.WriteFile(output, b, 0666)
}
