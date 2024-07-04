package main
import "fmt"
// 类型断言失败
func type_error() {
    var data interface{} = "hello world"

    number := data.(int)
    fmt.Println("数字是:", number)
}




