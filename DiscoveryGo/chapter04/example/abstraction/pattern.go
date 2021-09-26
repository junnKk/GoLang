package main

import (
	"fmt"
	"math"
)

type Func func(float64) float64 // Func는 실수값 하나를 받아서 실수값 하나를 돌려준다
type Transform func(Func) Func // Transfrom은 Func를 받아서 Func를 돌려준다. 이것은 함수 하나를 다른 함수 하나로 변환하는 함수의 형태

const tolerance = 0.00001
const dx = 0.00001

func Square(x float64) float64 { // 제곱을 구하는 함수
	return x * x
}

func FixedPoint(f Func, firstGuess float64) float64 { // 함수는 어떤 함수 f를 계속해서 적용했을 때, 어떤 값으로 수혐하는 경우에 그 수렴 값을 찾는 함수
	closeEnough := func(v1, v2 float64) bool { // 두 수 v1,v2가 서로 tolerance이하로 가까워졌으면 참을 돌려준다.
		return math.Abs(v1-v2) < tolerance
	}
	var try Func 
	try = func(guess float64) float64 { // try 함수는 guess에 반복적으로 함수 f를 적용시키다가 그 변화가 충분히 작으면 수렵된 것으로 보고 그 값을 반환하고 종료하는 함수이다.
		next := f(guess)
		if closeEnough(guess, next) {
			return next
		}
		return try(next)
	}
	return try(firstGuess)
}

func FixedPointOfTransform(g Func, transform Transform, guess float64) float64 {// FixedPoint의 아이디어를 함수로 변환하여 적용시켜 수렴 값을 찾는 것이다. 
	return FixedPoint(transform(g), guess)
}

func Deriv(g Func) Func {
	return func(x float64) float64 {
		return (g(x+dx) - g(x)) / dx
	}
}

func NewtonTransform(g Func) Func {
	return func(x float64) float64 {
		return x - (g(x) / Deriv(g)(x))
	}
}

func Sqrt(x float64) float64 {
	return FixedPointOfTransform(func(y float64) float64 {
		return Square(y) - x
	}, NewtonTransform, 1.0)
}

func main() {
	fmt.Println(Sqrt(4))
}
