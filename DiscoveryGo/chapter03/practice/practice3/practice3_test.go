package practice3

import "fmt"

func ExampleBinarySearch(){
	slice := []string{"가","나","다","라","마",}
	fmt.Println(BinarySearch(slice,0,4,"가"))
	fmt.Println(BinarySearch(slice,0,4,"바"))

	// Output:
	// true
	// false
}
