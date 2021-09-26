package hangul

import "fmt"

func ExampleHasConsonantSuffix(){
	fmt.Println(HasConsonantSuffixs("Go 언어"))
	fmt.Println(HasConsonantSuffixs("그럼"))
	fmt.Println(HasConsonantSuffixs("우리 밥 먹고 합시다."))
	// OutPut:
	// false
	// true
	// false
}

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()

	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("% x", s)
	// Output:
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	s := "가나다"
	b := []byte(s)
	b[2]++
	fmt.Println(string(b))

	// Output:
	// 각나다
}

func Example_printCompoundString() {
	s := "abcde가나다日本語"
	for _, c := range s {
		fmt.Printf("%c", c)
	}

	// Output:
	// abcde가나다日本語
}