package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	var s = "/asdf/dsafasfa/sadf/"
	join := filepath.Join("/", "/sdafaf/asdfadsf/dasfas/")
	fmt.Println(filepath.FromSlash(s))
	fmt.Println(filepath.FromSlash(join))

}
