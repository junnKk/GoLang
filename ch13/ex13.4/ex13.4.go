package main

import "fmt"

type User struct { // 일반 고객용 구조체
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct { // VIP 고객용 구조체
	User
	Level int
	Price int
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 10},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이 %d VIP 레벨: %d 유저 레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,
		vip.User.Level,
	)
}
