package main

import "fmt"
import "math"

func main(){
     var x, y int = 3, 4
     var f float64 = math.Sqrt(float64(x*x + y*y))
     var z uint = uint(f)
     fmt.Printf("%v, %v, float %v uint %v\n", x, y, f, z)
}