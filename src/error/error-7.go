package main

import "fmt"

type areaError7 struct {
	err    string //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError7) Error() string {
	return e.err
}

func (e *areaError7) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError7) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError7{err, length, width}
	}
	return length * width, nil
}

// 使用结构类型上的方法提供有关错误的更多信息
// 编写一个计算矩形面积的程序。如果长度或宽度小于零, 此程序将打印错误。
func main()  {
	length, width := -10.1, 10.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError7); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}