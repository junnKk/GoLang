package distribute

// 파이프라인 패턴에서 알아본 IntPipe 형태의 함수를 받은 뒤에 n개로 분산처리하는 함수로 돌려주는 함수.
// 그러므로 이 함수는 팬아웃과 팬인을 모두 수행하는 함수
func Distribute(p IntPipe, n int) IntPipe{
	return func(in <-chan int) <-chan int {
		cs:= make([]<-chan int, n)
		for i:=0;i < n; i++{
			cs[i] = p(in)
		}
		reutn FanIn(cs...)
	}
}