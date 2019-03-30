package main

import "fmt"

type Me struct {
    X int
    Y int
}

func main() {
    v := Me{1,5}
    p :=&v
    p.X = 1e9
    fmt.Println(v)
}
