/* 정렬된 문자열 슬라이스가 있을 때 특정 문자열이 슬라이스에 있는지를 조사하는 함수를 작성해보자 */ 

package practice3

func BinarySearch(arr []string, low, high int, target string) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	if arr[mid] == target {
		return true
	} else if arr[mid] > target {
		return BinarySearch(arr, low, mid-1, target)
	} else {
		return BinarySearch(arr, mid+1, high, target)
	}

}
