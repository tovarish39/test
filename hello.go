package main
 
import "fmt"


// var a string = "4"
func square() func() int { 
    var x int = 2
    return func() int { 
        fmt.Println("x =", x)
        x++
        return x * x
    }
}
 
func main() {
    f := square()
    fmt.Println(f())        // 9
    fmt.Println(f())        // 16
    fmt.Println(f())        // 25
}