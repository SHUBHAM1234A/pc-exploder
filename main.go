package main
import (
    "fmt"
    "strings"
)
func main() {
    a := []string{}
    a = append(a, "h")
    for i := 0; i <= 1000000000000000000000000000000000000; i++ {
        a = append(a, "m")
    }
    fmt.Println(strings.Join(a, ""))
}
