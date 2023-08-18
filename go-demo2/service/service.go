package service

import "go-demo2/dao"

type Serv struct {

}

func (s *Serv) Create(mod dao.User) error {

	return dao.Create(&mod)
}
