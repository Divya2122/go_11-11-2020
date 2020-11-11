package main

import (
	"fmt"
	"sort"
)

func main() {
	//sort slice containing strings
	strs := []string{"New York", "London", "Seattle"}
    sort.Strings(strs)
	fmt.Println("Strings:", strs)
	
	//sort slice containing ints
	ints := []int{17, 2, 14, 1, 5}
    sort.Ints(ints)
	fmt.Println("Ints: ", ints)
	
	//sort slice containing floats
	floats := []float64{23.93, 2.3, 2.5, 5.2}
	sort.Float64s(floats)
	fmt.Println("Floats: ", floats)
}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
package main
 
import (
	"fmt"
	"sort"
)
 
func main() {
	
	strs := []string{"New York", "London", "Seattle"}
    sort.Strings(strs)
	fmt.Println("Strings:", strs)
	
	
	ints := []int{17, 2, 14, 1, 5}
    sort.Ints(ints)
	fmt.Println("Ints: ", ints)
	

	floats := []float64{23.93, 2.3, 2.5, 5.2}
	sort.Float64s(floats)
	fmt.Println("Floats: ", floats)
}