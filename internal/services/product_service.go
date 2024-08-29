package services

import (
	"shopping-list-backend/internal/models"
	"shopping-list-backend/pkg/database"

	_ "github.com/lib/pq"
)

func GetProducts(listID int) ([]models.Product, error) {
	db, err := database.Connection()
	if err != nil {
		return nil, err
	}

	var products []models.Product

	query := `
		SELECT id, name, amount, completed, notes
		FROM products
		WHERE list_id = $1
		`

	rows, err := db.Query(query, listID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Amount, &product.Completed, &product.Notes)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func AddProductToList(listID int, product *models.Product) (int, error) {
	db, err := database.Connection()
	if err != nil {
		return -1, err
	}

	var lastInsertID int

	query := `
		INSERT INTO products (list_id, name, amount, completed, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`

	if product.Completed == nil {
		product.Completed = new(bool)
		*product.Completed = false
	}

	err = db.QueryRow(query, listID, product.Name, product.Amount, product.Completed, product.Notes).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil
}

func RemoveProduct(productID int) error {
	db, err := database.Connection()
	if err != nil {
		return err
	}

	query := `
		DELETE FROM products
		WHERE id = $1
		`

	_, err = db.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(product *models.Product) error {
	db, err := database.Connection()
	if err != nil {
		return err
	}

	query := `
		UPDATE products
		SET name = $1, amount = $2, completed = $3, notes = $4
		WHERE id = $5
		`
	_, err = db.Exec(query, product.Name, product.Amount, product.Completed, product.Notes, product.ID)
	if err != nil {
		return err
	}

	return nil
}
