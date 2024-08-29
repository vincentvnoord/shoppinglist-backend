package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shopping-list-backend/internal/models"
	"shopping-list-backend/internal/services"
	"strconv"

	"github.com/gorilla/mux"
)

func ProductsGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strListID := vars["list_id"]

	listID, err := strconv.Atoi(strListID)
	if err != nil {
		http.Error(w, "Invalid list ID", http.StatusBadRequest)
		return
	}

	products, err := services.GetProducts(listID)
	if err != nil {
		http.Error(w, "Problem fetching list: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if products == nil {
		http.Error(w, "List not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}

func ProductPost(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		ListID  int            `json:"list_id"`
		Product models.Product `json:"product"`
	}

	newProduct := RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil || newProduct.ListID == 0 || newProduct.Product.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	fmt.Println("Adding product to list", newProduct.Product.Name)
	productID, err := services.AddProductToList(newProduct.ListID, &newProduct.Product)
	if err != nil {
		http.Error(w, "Problem creating product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Product created with ID:", productID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": productID,
	})
}

func ProductPut(w http.ResponseWriter, r *http.Request) {
	updatedProduct := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil || updatedProduct.ID == nil || updatedProduct.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = services.UpdateProduct(&updatedProduct)
	if err != nil {
		http.Error(w, "Problem updating product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}
