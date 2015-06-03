package main

import "fmt"

var outerVal int = 42

func main() {
  x := "Hello World"
  fmt.Println(x)
  
  x = "first"
  fmt.Println(x)
  
  x = "second"
  fmt.Println(x)
  
  x = x + " third"
  fmt.Println(x)
  
  var y string = "Hello"
  var z string = "Hello"
  fmt.Println(y == z)
  
  foo := 5
  bar := 9
  fmt.Println("foo + bar: ", foo+bar)
  
  fmt.Println("outerVal should be 42: ")
  fmt.Println(outerVal)
  
  const thisIsConstant string = "this is a constant"
  fmt.Println(thisIsConstant)
  // thisIsConstant = "foobar" // error
  
  var (
    a = 5
    b = 10
    c = 15
  )
  fmt.Println("a+b+c: ", a+b+c)
  fmt.Println("-------------")
  
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)
  
  output := input * 2
  fmt.Println(output)
}