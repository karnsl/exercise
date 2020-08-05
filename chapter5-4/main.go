package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	answer := r1.Intn(100) + 1

	//fmt.Println("answer is ", answer)

	var input int
	var check string
	var i int
	for i = 0 ; i < 5 ; i++ {
		fmt.Printf("Please enter your number: ")
		fmt.Scanln(&input)
		if input < answer {
			check = "น้อยไป"
		} else if input > answer {
			check = "มากไป"
		} else {
			i = 99
			check = "เจอแล้ว"
		}
		fmt.Printf("%v\n", check)
	}
	if i == 5 {
		fmt.Println("เกินพอ")
	}
}