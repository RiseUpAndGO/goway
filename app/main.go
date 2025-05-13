package main

import "fmt"

type Сar struct {
	Name  string
	Speed int
}

func main() {
	m := make(map[string]Сar)
	m["Hyundai"] = Сar{"Santa Fe", 55}
	fmt.Println(m["Hyundai"].Name)
}
