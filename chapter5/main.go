package main

import "fmt"

func main(){
  // while style
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i = i + 1
  }
  
  fmt.Println("-------")
  
  // for loop style
  for i := 1; i <=10; i++ {
    fmt.Print(i)
    fmt.Print(": ")
    if i % 2 == 0 {
      fmt.Println("even")
    } else {
      fmt.Println("odd")
    }
  }
  
  fmt.Println("---------")
  
  for i := 1; i <=6; i++ {
    fmt.Print(i)
    fmt.Print(": ")
    switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("unknown number")
    }
  }
  
}