package main

import "fmt"

func hanio(n int, x, y, z string) {
	if n < 1 {
		return
	} else if n == 1 {
		fmt.Printf("%s -> %s \n", x, z)
		return
	} else {
		hanio(n-1, x, z, y)
		fmt.Printf("%s -> %s \n", x, z)
		hanio(n-1, y, x, z)
	}

}

func main() {
	hanio(20, "x", "y", "z")
}
