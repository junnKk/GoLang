/* 같은 원소가 여러 번 들어갈 수 있는 Multiset을 기본적으로 제공하는 맵을 이용하여 만들어보자 */

package practice5

import "strings"

// 새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() map[string]int{
	m := map[string]int{}
	return m
}

// Insert 함수는 집합에 val을 추가한다.
func Insert(m map[string]int, val string){
	m[val]++
}

// Erase 함수는 집합에서 val을 제거한다. 집합에 val이 없는 경우에는 아무 일도 일어나지 않는다. 
func Erase(m map[string]int, val string){
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

// Count 함수는 집합에 val이 들어있는 횟수를 구한다.
func Count(m map[string]int, val string) int{
	return m[val]
}

// String 함수는 집합에 들어있는 원소들을 {} 안에 빈 칸으로 구분하여 넣은 문자열을 반환한다. 
func String(m map[string]int) string{
	s := "{ "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}
