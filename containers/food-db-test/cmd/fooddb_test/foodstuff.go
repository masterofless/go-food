package main

import "fmt"
import "os"

type Foodstuff struct {
    name string
}

func newFoodstuff() Foodstuff {
    f := Foodstuff{"Brussel Sprouts"}
    return f
}

func (f *Foodstuff) ToString() string {
    return "My Foodname is " + f.name
}

func main() {
    f := newFoodstuff()
    fmt.Fprintf(os.Stdout, "Starting Foodstuff %v", f)
}
