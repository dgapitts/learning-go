## time package and common standard time formats 



Here are some examples of common standard time formats as per https://www.geeksforgeeks.org/time-formatting-in-golang/
```
Format      Example
ANSIC	    “Mon Jan _2 15:04:05 2006”
UnixDate	“Mon Jan _2 15:04:05 MST 2006”
RubyDate	“Mon Jan 02 15:04:05 -0700 2006”
RFC822	    “02 Jan 06 15:04 MST”
RFC822Z	    “02 Jan 06 15:04 -0700”
RFC850	    “Monday, 02-Jan-06 15:04:05 MST”
RFC1123	    “Mon, 02 Jan 2006 15:04:05 MST”
RFC1123Z	“Mon, 02 Jan 2006 15:04:05 -0700”
RFC3339	    “2006-01-02T15:04:05Z07:00”
RFC3339Nano	“2006-01-02T15:04:05.999999999Z07:00”
```

I liked the usage of `` from https://gobyexample.com/time-formatting-parsing


```
~/projects/learning-go/ex004_time_pkg $ cat time.go
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
...
```
which returns 

```
~/projects/learning-go/ex004_time_pkg $ go run .
2022-01-12T21:31:05+01:00
...
```

and the rest of the code here is from https://www.geeksforgeeks.org/time-formatting-in-golang/

```
    // https://www.geeksforgeeks.org/time-formatting-in-golang/
    current_time := time.Now()
    fmt.Println("ANSIC: ", current_time.Format(time.ANSIC))
    fmt.Println("UnixDate: ", current_time.Format(time.UnixDate))
    fmt.Println("RFC1123: ", current_time.Format(time.RFC1123))
    fmt.Println("RFC3339Nano: ", current_time.Format(time.RFC3339Nano))
    fmt.Println("RubyDate: ", current_time.Format(time.RubyDate))
```

which returns 

```
~/projects/learning-go/ex004_time_pkg $ go run .
...
ANSIC:  Wed Jan 12 21:32:32 2022
UnixDate:  Wed Jan 12 21:32:32 CET 2022
RFC1123:  Wed, 12 Jan 2022 21:32:32 CET
RFC3339Nano:  2022-01-12T21:32:32.724611+01:00
RubyDate:  Wed Jan 12 21:32:32 +0100 2022
```
