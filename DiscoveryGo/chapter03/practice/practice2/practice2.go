/* []int 슬라이스를 넘겨받아서 오름차숭으로 정렬하는 함수를 작성하라 */ 
package practice2

import "fmt"

func SortInt() {
	nums := []int{1, 4, 3, 2, 5}
	for i := range nums {
		for j := i + 1; j<len(nums); j++ {
			if nums[i]>nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
}

