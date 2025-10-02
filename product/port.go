package product

import (
	"github.com/enghasib/server/domain"
	productHandler "github.com/enghasib/server/rest/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Create(product domain.Product) (*domain.Product, error)
	Get(productId int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(id int, product domain.Product) (*domain.Product, error)
	Delete(id int) error
}
