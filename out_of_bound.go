package main

import "fmt"
// out of bound
func out_of_bound() {
    numbers := []int{1, 2, 3}
    for i := 0; i <= len(numbers); i++ { 
        fmt.Println(numbers[i]) 
    }
}