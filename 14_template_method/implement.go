package templatemethod

import (
	"errors"
	"fmt"
)

type User struct {
	name     string
	password string
}

// 通用接口
type Dao interface {
	Add(object interface{}) error
}

// 子类具体实现接口
type service interface {
	valid(object interface{}) error
	create(object interface{}) error

	add(object interface{}) error
	update(object interface{}) error
	query(object interface{}) error
	delete(object interface{}) error
}

type daoService struct {
	service
}

func newService(service service) *daoService {
	return &daoService{
		service: service,
	}
}

func (d *daoService) Add(object interface{}) error {
	if err := d.valid(object); err != nil {
		return err
	}

	if err := d.create(object); err != nil {
		return err
	}

	if err := d.add(object); err != nil {
		return err
	}

	fmt.Println("add object success")
	return nil
}

type UserDaoService struct {
	*daoService
}

func NewUserDaoService() Dao {
	userDaoService := &UserDaoService{}
	daoService := newService(userDaoService)
	userDaoService.daoService = daoService
	return userDaoService
}

func (UserDaoService) valid(object interface{}) error {
	fmt.Println("valid")
	user, ok := object.(*User)
	if !ok {
		return errors.New("object not user")
	}
	if len(user.name) <= 0 {
		return errors.New("valid username error")
	}
	if len(user.password) <= 0 {
		return errors.New("valid user password error")
	}
	return nil
}

func (u *UserDaoService) create(object interface{}) error {
	fmt.Println("create")
	return nil
}

func (u *UserDaoService) add(object interface{}) error {
	fmt.Println("add")
	return nil
}

func (u *UserDaoService) update(object interface{}) error {
	return nil
}

func (u *UserDaoService) delete(object interface{}) error {
	return nil
}

func (u *UserDaoService) query(object interface{}) error {
	return nil
}
