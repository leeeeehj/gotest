package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z:=2.0
    diff:=1.0
    for diff>0.00000001 {
        val1:=z
        z=z-(z*z-x)/(2*z)
        diff=val1-z

    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
