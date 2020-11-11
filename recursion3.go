package main  
import (  
   "fmt"  
)  
func main() {  
   fmt.Println(factorial(1))  
}  
func factorial(num int ) int{  
   if num == 1{  
      return 0  
   }  
   return num*factorial(num-1)  
}  