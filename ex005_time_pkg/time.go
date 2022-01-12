package main

import (
    "fmt"
    "time"
)

func main() {

    // https://gobyexample.com/time-formatting-parsing
    p := fmt.Println
    t := time.Now()
    p(t.Format(time.RFC3339))

    // https://www.geeksforgeeks.org/time-formatting-in-golang/
    current_time := time.Now()
    fmt.Println("ANSIC: ", current_time.Format(time.ANSIC))
    fmt.Println("UnixDate: ", current_time.Format(time.UnixDate))
    fmt.Println("RFC1123: ", current_time.Format(time.RFC1123))
    fmt.Println("RFC3339Nano: ", current_time.Format(time.RFC3339Nano))
    fmt.Println("RubyDate: ", current_time.Format(time.RubyDate))

}


