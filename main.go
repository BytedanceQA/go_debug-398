package main

import (
    "fmt"
    "os"
)


func main() {
    // 定义一个映射，其中键是字符串，值是对应的函数
    functions := map[string]func(){
        "concurrency_error": concurrency_error,
        "out_of_bound":    out_of_bound,
		"nil_poniter":  nil_poniter,
		"type_error" : type_error}

    if len(os.Args) < 2 {
        fmt.Println("请提供一个函数名作为参数")
        return
    }

    // 从映射中查找函数
    if fn, ok := functions[os.Args[1]]; ok {
        fn() // 调用函数
    } else {
        fmt.Println("未知的函数名")
    }
}

