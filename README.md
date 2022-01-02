
## helloworld in go

We need to declare a main package and import fmt but still a very simple example:

```
~/projects/learning-go $ cat helloworld.go
package main

import "fmt"

func main() {
   fmt.Println("Hello World")
}
```
and to run this 
```
~/projects/learning-go $ go run helloworld.go
Hello World
```


