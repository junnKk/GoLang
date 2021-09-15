package slice

import "fmt"

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])

	n := 1
	m := 3
	fmt.Println(nums[n:m])

	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
	// [2 3]
}
func Example_append() {
	var fruits []string
	fmt.Println(fruits)

	fruits = append(fruits, "grape")
	fmt.Println(fruits)

	fruits2 := []string{"strawberry", "banana", "tomato"}
	fruits3 := append(fruits, fruits2[:2]...)
	fmt.Println(fruits3)

	// Output:
	// []
	// [grape]
	// [grape strawberry banana]
}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len: ", len(nums))
	fmt.Println("cap: ", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len: ", len(sliced1))
	fmt.Println("cap: ", cap(sliced1))
	fmt.Println()

	sliced2 := nums[:2]
	fmt.Println(sliced2)
	fmt.Println("len: ", len(sliced2))
	fmt.Println("cap: ", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len: ", len(sliced3))
	fmt.Println("cap: ", cap(sliced3))
	fmt.Println()

	sliced3[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
	fmt.Println()

	nums2 := make([]int, 0, 1)
	fmt.Println(len(nums2), cap(nums2))

	nums2 = append(nums2, 1)
	nums2Location := &nums2[0]
	fmt.Println(len(nums2), cap(nums2))

	nums2 = append(nums2, 2)
	nums2Location2 := &nums2[0] // slice nums2 should be copied into another memory location
	fmt.Println(len(nums2), cap(nums2))

	fmt.Println(nums2Location == nums2Location2)

	// Output:
	// 	[1 2 3 4 5]
	// len:  5
	// cap:  5
	//
	// [1 2 3]
	// len:  3
	// cap:  5
	//
	// [1 2]
	// len:  2
	// cap:  5
	//
	// [1 2 3 4]
	// len:  4
	// cap:  5
	//
	// [1 2 100 4 5] [1 2 100] [1 2] [1 2 100 4]
	//
	// 0 1
	// 1 1
	// 2 2
	// false
}

func Example_sliceCopy() {
	// example for deepcopy in python

	// 1. naive solution
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)

	// 2. built-in function
	// 2-1. check copy function
	dest2 := make([]int, len(src)-1)
	if n := copy(dest2, src); n != len(src) {
		fmt.Println("copy is not completed")
	}
	// 2-2. right usage
	dest3 := make([]int, len(src))
	copy(dest3, src)
	fmt.Println(dest3)

	// 3. using append function
	dest4 := append([]int(nil), src...)
	fmt.Println(dest4)

	// Output:
	// [1 2 3 4 5]
	// copy is not completed
	// [1 2 3 4 5]
	// [1 2 3 4 5]
}

func Example_sliceTrick() {
	// futher reference
	// https://github.com/golang/go/wiki/SliceTricks

	a := []int{1, 2, 3, 4, 5}

	// 1. Insert
	// 1-1 Insert single element
	insertedNum := 10
	i := 3
	a = append(a, 0)
	copy(a[i+1:], a[i:])
	a[i] = insertedNum
	fmt.Println(a)
	// 1-2 Insert multiple element
	insertedSlice := []int{4, 5, 6}
	a = append(a, insertedSlice...)
	copy(a[i+len(insertedSlice):], a[i:])
	copy(a[i:], insertedSlice)
	fmt.Println(a)

	// 2. Delete with preserve order
	// 2-1 Delete single element
	i = 4
	a = a[:i+copy(a[i:], a[i+1:])]
	fmt.Println(a)
	// 2-2 Delete multiple element: O(n)
	i = len(a) - 3
	k := 5
	if i+k > len(a) {
		k = len(a) - i
	}
	a = a[:i+copy(a[i:], a[i+k:])]
	fmt.Println(a)

	// 3. Delete without preserve order: O(k)
	// 2-1 Delete single element
	i = 3
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]
	fmt.Println(a)
	// 2-1 Delete multiple element
	i = 1
	k = 2
	start := len(a) - k
	if i+k > start {
		start = i + k
	}
	copy(a[i:i+k], a[start:])
	a = a[:len(a)-k]
	fmt.Println(a)

	// Output:
}
