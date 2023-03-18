
# 前言

Go语言内置了 net/http库，封装了HTTP网络编程的基础的接口，一个http服务器就可以正常运行访问接受请求

```go
package main
import (
    "fmt"
    "net/http"
)
func myfunc(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hi")
}

func main() {
    http.**HandleFunc**("/", myfunc)
    http.ListenAndServe(":8080", nil)
}
```

下面深入了解一下过程


****