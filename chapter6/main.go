package main

import "fmt"

func main(){
  var x [5]int
  x[4] = 100
  fmt.Println(x)
  fmt.Println(x[4])
  
  getAverage()
}

func getAverage(){
  
  vals := [5]float64{ 98, 93, 77, 82, 83 }
  
  var total float64 = 0
  
  arrLen := len(vals)
  
  /*
  for i := 0; i < arrLen; i++ {
    total += x[i]
  }
  */
  
  // this loop may take time getting used to:
  for _, value := range vals {
    total += value
  }
  
  fmt.Println(total / float64(arrLen))
}