package main

import "fmt"

func fibonaci(i int) (ret int) {
   if i == 0 {
      return 0
   }
   if i == 1 {
      return 1
   }
   return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
   var i int
   for i = 1; i < 15; i++ {
      fmt.Printf("%d ", fibonaci(i))
   }
}