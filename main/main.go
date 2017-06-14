package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/pallat/godx"
)

var file string
var output string
var pack string

func init() {
	flag.StringVar(&file, "f", "example.xml", "-f=example.xml")
	flag.StringVar(&output, "o", "example.go", "-o=example.go")
	flag.StringVar(&pack, "p", "main", "-p=packageName")
}

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	x := godx.New(b)
	ioutil.WriteFile(output, []byte(x.MakePackage(pack)), 0666)

	// b, err = format.Source([]byte(x.String()))

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// ioutil.WriteFile(output, b, 0666)
}
