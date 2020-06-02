package main

import "fmt"

func main () {
     names := [4]string{
     	   "Jose",
	   "David",
	   "Caro",
	   "Cantor",
     }

     fmt.Println(names)
     a := names[0: 3]
     b := names[3: ]
     names[0] = "Josese"
     a[1] = "Davidid"
     fmt.Println(a, b)
     fmt.Println(names)
}