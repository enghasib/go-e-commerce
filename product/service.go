package product

import "github.com/enghasib/server/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (srv *service) Create(product domain.Product) (*domain.Product, error) {
	createdProduct, err := srv.productRepo.Create(product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}

func (srv *service) Get(productId int) (*domain.Product, error) {
	product, err := srv.productRepo.Get(productId)
	if err != nil {
		return nil, err
	}
	return product, err
}

func (srv *service) List() ([]*domain.Product, error) {
	listOfProduct, err := srv.productRepo.List()
	if err != nil {
		return nil, err
	}
	return listOfProduct, nil
}

func (srv *service) Update(id int, product domain.Product) (*domain.Product, error) {
	prod, err := srv.productRepo.Update(id, product)
	if err != nil {
		return nil, err
	}
	return prod, err
}

func (srv *service) Delete(id int) error {
	if err := srv.productRepo.Delete(id); err != nil {
		return err
	}
	return nil
}
