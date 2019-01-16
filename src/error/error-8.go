package main

import (
	"fmt"
	"errors"
)

type areaError8 struct {
	err    string //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError8) Error() string {
	return e.err
}

func (e *areaError8) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError8) widthNegative() bool {
	return e.width < 0
}

var lengthError = errors.New("length is less than zero")
var widthError = errors.New("width is less than zero")

func rectArea2(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
		return 0, lengthError
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
		return 0, widthError
	}
	if err != "" {
		return 0, &areaError8{err, length, width}
	}
	return length * width, nil
}

// 使用直接比较的方法提供有关错误的更多信息
// 编写一个计算矩形面积的程序。如果长度或宽度小于零, 此程序将打印错误。
func main()  {
	length, width := -10.1, -10.0
	area, err := rectArea2(length, width)
	if err != nil {
		if err == lengthError {
			fmt.Printf("error: length %0.2f is less than zero\n")
		} else if err == widthError {
			fmt.Printf("error: width %0.2f is less than zero\n")
		} else {
			fmt.Println(err)
		}
		return
	}
	fmt.Println("area of rect", area)
}