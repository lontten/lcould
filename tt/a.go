package main

import "fmt"

func main() {
	a := map[string]int{}
	i, ok := a["a"]
	fmt.Println(i, ok)

}
