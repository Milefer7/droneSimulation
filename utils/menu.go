package utils

import "fmt"

// DisplayMenu 显示菜单并获取用户选择
func DisplayMenu() int {
	fmt.Println("\n菜单:")
	fmt.Println("1. 运行仿真")
	fmt.Println("2. 退出")
	fmt.Print("请选择: ")
	var choice int
	fmt.Scan(&choice)
	return choice
}
