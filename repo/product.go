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

func (r *productRepo) List() ([]*domain.Product, error) {
	var productList []*domain.Product
	query := `
		SELECT id, title, description, price, img_url FROM products
	`
	err := r.db.Select(&productList, query)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return productList, nil
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

// func generateProduct(r *productRepo) {
// 	var ProductList = []*Product{
// 		{ID: 1, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgUrl: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},

// 		{ID: 2, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgUrl: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},

// 		{ID: 3, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},

// 		{ID: 4, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgUrl: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"},

// 		{ID: 5, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgUrl: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528"},

// 		{ID: 6, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s"},
// 	}
// 	r.productList = append(r.productList, ProductList...)

// }
