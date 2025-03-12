package main

import (
	"fmt"
)

// 讓使用者輸入一個字元，並判斷字元是大寫、小寫、數字或其他字元
func main() {
	var input string
	fmt.Scanln(&input)

	if len(input) > 0 {
		char := input[0]

		if char >= 'A' && char <= 'Z' {
			fmt.Println("大寫")
		} else if char > 'a' && char < 'z' {
			fmt.Println("小寫")
		} else if char >= '0' && char <= '9' {
			fmt.Println("數字")
		} else {
			fmt.Println("其它字元")
		}
	}
}
