package main
import "fmt"

type Me struct {
    X int
    Y int
}

func main() {
   v := Me{1,2}
   fmt.Println(v.X)
   v.X = 4
   fmt.Println(v.X)
}
