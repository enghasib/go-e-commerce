package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/enghasib/server/domain"
	"github.com/enghasib/server/product"
	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	// productList []*Product,
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	// repo := &productRepo{}
	// generateProduct(repo)
	// return repo
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(product domain.Product) (*domain.Product, error) {

	query := `
		INSERT INTO products (title, description, price, img_url) values(
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id
	`
	rows := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImgUrl)
	if rows.Err() != nil {
		fmt.Println("error:", rows.Err())
		return nil, rows.Err()
	}
	rows.Scan(&product.ID)
	return &product, nil
}

func (r *productRepo) Get(productId int) (*domain.Product, error) {
	var prod domain.Product
	query := `
		SELECT id, title, description, price, img_url FROM products WHERE id=$1
	`
	err := r.db.Get(&prod, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println("error:", err)
		return nil, err
	}
	return &prod, nil
}

func (r *productRepo) List(limit, page int) ([]*domain.Product, error) {

	offset := ((page - 1) * limit) + 1

	var productList []*domain.Product
	query := `
		SELECT id, title, description, price, img_url FROM products LIMIT $1 OFFSET $2;
	`
	err := r.db.Select(&productList, query, limit, offset)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return productList, nil
}

func (r *productRepo) Count() (int, error) {
	query := `
		SELECT COUNT(*) FROM products;
	`
	var count int
	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {
		fmt.Println("error:", err)
		return 0, err
	}
	return count, nil
}

func (r *productRepo) Update(prodId int, product domain.Product) (*domain.Product, error) {
	// for i := range p.productList {
	// 	if p.productList[i].ID == id {
	// 		if &product.Title != nil {
	// 			p.productList[i].Title = product.Title
	// 		}
	// 		if &product.Description != nil {
	// 			p.productList[i].Description = product.Description
	// 		}
	// 		if &product.Price != nil {
	// 			p.productList[i].Price = product.Price
	// 		}
	// 		if &product.ImgUrl != nil {
	// 			p.productList[i].ImgUrl = product.ImgUrl
	// 		}
	// 	}
	// 	return p.productList[i], nil
	// }
	// return nil, errors.New("update failed!")
	var existingProduct domain.Product
	// fetch product
	err := r.db.Get(&existingProduct, `select id, title, description, price, img_url from products where id=$1`, prodId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Product not found")
			return nil, nil
		}
		fmt.Println("error:", err)
		return nil, err
	}
	// check if product is already exist
	if product.Title != "" {
		existingProduct.Title = product.Title
	}
	if product.Description != "" {
		existingProduct.Description = product.Description
	}
	if product.Price != 0 {
		existingProduct.Price = product.Price
	}
	if product.ImgUrl != "" {
		existingProduct.ImgUrl = product.ImgUrl
	}

	query := `
		UPDATE products
		SET 
		title=$1, description=$2,
		price=$3, img_url=$4
		WHERE id=$5
		RETURNING id, title, description, price, img_url
	`
	err = r.db.Get(
		&existingProduct,
		query,
		product.Title,
		product.Description,
		product.Price,
		product.ImgUrl,
		prodId,
	)

	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return &existingProduct, nil
}

func (r *productRepo) Delete(productId int) error {
	// index := -1

	// for i, product := range p.productList {
	// 	if product.ID == productId {
	// 		index = i
	// 		break
	// 	}
	// }
	// p.productList = append(p.productList[:index], p.productList[index+1:]...)
	// return nil

	query := `
		DELETE FROM products WHERE id = $1 
	`

	result, err := r.db.Exec(query, productId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowEffected, _ := result.RowsAffected()
	if rowEffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
