package main
 
import (
	"fmt"
	"sort"
)
 
type Emp struct {
	FirstName string
	LastName string
	Age int
}
 
type ByAge []Emp
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
 
func main() {
	emps := []Emp{
        {"Rahul", "Anand", 29},
        {"C", "Anderson" , 20},
        {"Bob", "Anderson", 25},
	}
	
	sort.Sort(ByAge(emps))
	fmt.Println(emps)
}