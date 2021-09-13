// 필드 배치 순서에 따른 구조체 크기 변화

package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age int32 // 4바이트
	Score float64 // 8바이트
}

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
	// unsafe.Sizeof()는 해당 변수의 메모리 공간 크기를 반환.
	// 앞에서 배운대로라면 이번 구조체의 User 크기는 12 바이트여야 하는데 16바이트로 출력되는 이유는 메모리 정렬 때문

}


// 메모리 정렬: 컴튜터가 데이터에 효과적으로 접근하고자 메모리를 일정 크기 간격으로 정렬하는 것
// 메로리 패딩: 메모리 정렬을 위해서 필드 사이에 공간을 띄우는 것
//			- 4바이트 변수의 시작 주소는 4의 배수로 맞추고 2바이트 변수의 시작 주소는 2의 배수로 맞춰서 패딩함.
