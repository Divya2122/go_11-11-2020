package main
 
import (
	"fmt"
	"sort"
)
 
func main() {
	emps := []struct {
		FirstName string
		LastName string
		Age int
	}{
        {"Rahul", "Anand", 29},
        {"C", "Anderson" , 20},
        {"Bob", "Anderson", 25},
	}
	
	sort.Slice(emps, func(i, j int) bool {
		return emps[i].Age < emps[j].Age
	})
 
	fmt.Println(emps)
}