package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"product-rest/model"
	"strings"
	"time"
)

type ProductRepository interface{
	GetAllProducts(sortBy string, typeSort string) ([]model.Product, error)
	InsertProduct(product model.Product) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db: db}
}

func(r *productRepository) GetAllProducts(sortBy string, typeSort string) ([]model.Product, error){

	query := "SELECT name, price, description, quantity, created_at FROM product"

	if sortBy != "" {
		orderBy := fmt.Sprintf(" ORDER BY %s", sortBy)
		if typeSort != "" {
			orderBy = fmt.Sprintf("%s %s", orderBy, strings.ToUpper(typeSort))
		}

		query += orderBy
	}

	result, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	products := []model.Product{}
	
	for result.Next() {

		var item model.Product
		var createdAt string
		
		err := result.Scan(&item.Name, &item.Price, &item.Description, &item.Quantity, &createdAt)
		if err != nil {
			return nil, err
		}

		item.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}

		products = append(products, item)
	}

	return products, nil
}

func(r *productRepository) InsertProduct(product model.Product) error {

	res, err := r.db.Exec("INSERT INTO product(name, price, description, quantity) VALUE(?,?,?,?)", product.Name, product.Price, product.Description, product.Quantity)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("failed insert products")
	}

	return nil
}