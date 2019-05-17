package main

import "fmt"

func main() {
	const s = "squk"
	for i := 1; ; i = (i + 1) % 2 {
		fmt.Printf("You're a %sid now!\n", s[i*3:3+i])
	}
}
