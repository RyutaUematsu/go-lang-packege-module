package sample_pkg_1

import "fmt"

// export するときは 大文字にする
func Hello(name string) string {
	var introduction string
	introduction = fmt.Sprintf("hello %s", name)
	return introduction
}
