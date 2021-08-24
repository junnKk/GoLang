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


}
