package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shopping-list-backend/internal/models"
	"shopping-list-backend/internal/services"

	"github.com/gorilla/mux"
)

func ListGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListGet")
	vars := mux.Vars(r)
	publicCode := vars["public_code"]

	fmt.Println("Public code:", publicCode)

	list, err := services.GetListByCode(publicCode)
	if err != nil {
		http.Error(w, "Problem fetching list: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if list == nil {
		http.Error(w, "List not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(list)
}

func ListPost(w http.ResponseWriter, r *http.Request) {
	newList := models.List{}
	err := json.NewDecoder(r.Body).Decode(&newList)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	listID, publicCode, err := services.CreateList(&newList)
	if err != nil {
		http.Error(w, "Problem creating list: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("List created with ID:", listID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":          listID,
		"public_code": publicCode,
	})
}

func ListPut(w http.ResponseWriter, r *http.Request) {

}
