package test

import (
	. "shopping-list-backend/internal/models"
	"shopping-list-backend/internal/services"
	"shopping-list-backend/pkg/database"
	"testing"
)

func TestAddRemoveProduct(t *testing.T) {
	database.LoadEnv()

	firstList, err := services.GetFirstList()
	if err != nil {
		t.Errorf("Error getting first list, %v", err)
	}

	completed := false
	notes := "This is a note"

	product := Product{
		Name:      "Product 1",
		Completed: &completed,
		Notes:     &notes,
	}

	productID, err := services.AddProductToList(firstList, &product)

	if err != nil {
		t.Errorf("Error adding product to list, %v", err)
	}

	t.Log(productID)

	//err = services.RemoveProduct(productID)
	if err != nil {
		t.Errorf("Error removing product, %v", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	database.LoadEnv()

	firstList, err := services.GetFirstList()
	if err != nil {
		t.Errorf("Error getting first list, %v", err)
	}

	amount := "100"
	notes := "This is a note"
	completed := false

	product := Product{
		Name:      "Product 1",
		Amount:    &amount,
		Completed: &completed,
		Notes:     &notes,
	}

	productID, err := services.AddProductToList(firstList, &product)

	if err != nil {
		t.Errorf("Error adding product to list, %v", err)
	}

	product.Name = "Updated product"
	err = services.UpdateProduct(&product)
	if err != nil {
		t.Errorf("Error updating product, %v", err)
	}

	err = services.RemoveProduct(productID)
	if err != nil {
		t.Errorf("Error removing product, %v", err)
	}
}
