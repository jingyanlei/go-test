package main

import (
	"fmt"
	"math"
)

// 结构类型有一个字段radius , 它存储负责错误的半径值, 而err字段存储实际的错误消息。
type areaError struct {
	err		string
	radius	float64
}

// 实现错误接口
func (e areaError)Error() string {
	return fmt.Sprintf("radius %0.2f %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"Radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}
