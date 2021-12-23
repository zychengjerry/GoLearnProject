package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var mark int = 90

	switch mark {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("秀\n")
	case grade == "B":
		fmt.Printf("良好\n")
	case grade == "C":
		fmt.Printf("加油\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不合格\n")
	default:
		fmt.Printf("寄\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
