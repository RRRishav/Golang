package main

import "fmt"

func main() {
	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	
    case 2:
	    fmt.Println("two")

	default:
		fmt.Println("no case")
}
}