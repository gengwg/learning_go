package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {

    // var s string
    // fmt.Scanln(&s)
    // fmt.Println(s)

    // enter string
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    str, _ := reader.ReadString('\n')
    fmt.Println(str)

    // enter number
    fmt.Print("Enter a number: ")
    str, _ = reader.ReadString('\n')
    // need trim string \n before and after
    f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Value of number:", f)
    }
}
