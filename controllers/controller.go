package controllers

import "gorm.io/gorm"

/*
Controller interface:
- Controller has a generic type Entity
- controllers support the following methods:
	- GetAll()
	- Get(id int)
	- Create(element Entity)
	- Update(id int, element Entity)
	- Delete(id int)
 */

type Controller struct {
	DB *gorm.DB
}

type IController[E any] interface {
	GetAll() ([]E, error)
	Get(id int) (E, error)
	Create(element E) (E, error)
	Update(id int, element E) (E, error)
	Delete(id int) error
}
