package main

import (
	"flag"
	"fmt"
)

func main() {
	//
	name := flag.String("name", "value", "specify the name you want to say it")
	fmt.Println(*name)
}
