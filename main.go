package main

import (
	"fmt"
	"test_module/sample_pkgs/sample_pkg_1"
	"test_module/sample_pkgs/sample_pkg_2"
)

func main() {
	var name string = "Ryuta"
	hello := sample_pkg_1.Hello(name)
	fmt.Println("Hello Go lang")
	fmt.Println(hello)
	fmt.Println(sample_pkg_2.Func1())
	fmt.Println(sample_pkg_2.Func2())
}
