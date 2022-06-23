package main

import (
	"fmt"
	"runoob/circle"
)

func main() {
	var radius float64
	print("请输入圆的半径(米): ")
	_, err := fmt.Scan(&radius)
	if err != nil {
		return
	}
	c := circle.NewCircle(radius)
	println(c.ToString())
}
