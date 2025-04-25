package models

import "fmt"

type UserModel struct {
	Id   int
	Name string
}

func (u *UserModel) Key() int {
	return len(u.Name)
}

func (u *UserModel) ToString() string {
	return fmt.Sprintf("%d - %s", u.Id, u.Name)
}
