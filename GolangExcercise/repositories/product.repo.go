package repositories

import "golangexcercise/minhnq4/entities"

type ProductRepo interface {
	GetAll() []*entities.Product
	GetById(id int) (*entities.Product, bool)
	Create(product *entities.Product) int
	DeleteById(id int)
	SaveData()
	GetData() 
}
