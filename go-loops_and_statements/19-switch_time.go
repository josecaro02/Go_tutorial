package main

import (
       "fmt"
       "time"
)

func main () {
     fmt.Println("When's Saturday?")
     today := time.Now().Weekday()
     fmt.Println(time.Saturday, " and ", today)
     switch time.Saturday {
     	    case today + 0:
	    	 fmt.Println("Today.")
	    case today + 1:
	    	 fmt.Println("Tomorrow.")
	    default:
		fmt.Println("Too far way")
     }
}