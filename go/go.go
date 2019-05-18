package main

import "fmt"

func main() {
	var j int
	const s = "ksqu"
	for i := 0; ; i = (i + 1) % 2 {
		if j++; j > 9 {
			return
		}
		fmt.Printf("You're a %sid now!\n", s[i:1+i*3])
	}
}
