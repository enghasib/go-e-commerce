package product

import "github.com/enghasib/server/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	List(limit, page int) ([]*domain.Product, error)
	Count() (int, error)
	Get(id int) (*domain.Product, error)
	Update(id int, prod domain.Product) (*domain.Product, error)
	Delete(id int) error
}
