package dao

import "fmt"

type Say interface {
	create() error
}

func Create(mod Say) error {
	return mod.create()
}

func (s *Studest) create() error {
	fmt.Println("create - studest ", s.Name)
	return nil
}

type Studest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) create() error {
	fmt.Println("create - user ", u.Name)
	return nil
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
