package main

import (
	"fmt"
	"math"
)

// 使用Errorf函数将更多信息添加到错误中
func circleArea5(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea5(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}