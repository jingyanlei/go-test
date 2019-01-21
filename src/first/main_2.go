package main

import "fmt"

// 定义学生类型
type student struct {
	firstName string
	lastName string
	grade string
	country string
}

// 编写filter函数。该函数采用一小部分学生和一个函数来确定学生是否将过滤标准与参数相匹配
func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)

	f2 := filter(s, func(s student) bool {
		if s.country == "India" {
			return true
		}
		return false
	})
	fmt.Println(f2)

}