package user

import "fmt"

type user struct {
	name	string
	sex		int
}

func New(name string, sex int) user {
	user := user{
		name:name,
		sex:sex,
	}
	return user
}

func (u user)GetUser()  {
	fmt.Printf("user:%s, sex:%d\n", u.name, u.sex)
}