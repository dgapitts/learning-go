
## Starting with helloworld in go - run direcrly or build an executable

We need to declare a main package and import fmt but still a very simple example:

```
~/projects/learning-go/ex001_helloworld $ cat helloworld.go
package main

import "fmt"

func main() {
   fmt.Println("Hello World")
}
```

to **run** this directly, either use `go run helloworld.go` or `go run .`
``` 
~/projects/learning-go/ex001_helloworld $ go run helloworld.go
Hello World
~/projects/learning-go/ex001_helloworld $ go run .
Hello World
```
and to **build** this as a MacOSX specific binary
```
~/projects/learning-go/ex001_helloworld $ ls -ltr
total 8
-rw-r--r--  1 dave  staff  83 Jan  2 21:12 helloworld.go
~/projects/learning-go/ex001_helloworld $ go build .
~/projects/learning-go/ex001_helloworld $ ls -ltr
total 3664
-rw-r--r--  1 dave  staff       83 Jan  2 21:12 helloworld.go
-rwxr-xr-x  1 dave  staff  1869504 Jan  3 17:35 ex001_helloworld
~/projects/learning-go/ex001_helloworld $ ./ex001_helloworld
Hello World
~/projects/learning-go/ex001_helloworld $ ls -ltrh
total 3664
-rw-r--r--  1 dave  staff    83B Jan  2 21:12 helloworld.go
-rwxr-xr-x  1 dave  staff   1.8M Jan  3 17:35 ex001_helloworld
```




