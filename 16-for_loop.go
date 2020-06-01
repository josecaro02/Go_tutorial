package main

import "fmt"

func main () {
     sum := 0
     for i := 0; i < 10; i++ {
	 sum += i
	 fmt.Printf("Value: %v and %v\n", sum, i)
     }
}