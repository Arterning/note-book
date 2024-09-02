
Go mod init and run and build

```bash
mkdir my-app
cd my-app
go mod init my-app

go run main.go
go build -o app.exe

# -u means update
go get -u github.com/gin-gonic/gin

# create main.go
go run main.go

# create lib.go and you use somethings in lib.go
go run main.go lib.go
go build
```


To invoke the `Sub` function from "lib.go" instead of defining another `Sub` function in "main.go", you can update the "main.go" file as follows:


lib.go
```go
package main

func Sub(a, b int) int {
    return a - b
}
```


main.go
```go
package main

import (
	"fmt"
)

func main() {
	result := Sub(3, 34)
	fmt.Println("Result:", result)
}
```

Make sure that "lib.go" and "main.go" are in the same package. This will allow the Sub function from "lib.go" to be invoked in "main.go".

That means if you put the lib.go and main.go in the same path, you just need to make the method to Uppercase, sub->Sub, then it is public. you can just use it without import(it is the same as java, if is the same pacage, import is not necessary)


when i run 'go build', it works , but if i run 'go run main.go' ,error:undefined: Sub

When you run go build, it successfully compiles the package, including the Sub function from "lib.go". However, when you run go run main.go, the main.go file is compiled and executed as a standalone file, and it does not automatically include the other files in the package.

To run the program using go run, you can specify all the files that need to be included in the command, like this:

```bash
go run main.go lib.go
```

we recommend you use
```
go build
```




If you put lib.go in another directory, like 'lib', and you use`package lib` at the top

you should import it

```go
import lib from 'modulename/lib'
```







