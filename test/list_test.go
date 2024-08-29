package test

import (
	. "shopping-list-backend/internal/models"
	"shopping-list-backend/internal/services"
	"shopping-list-backend/pkg/database"
	"testing"
)

func TestAddRemoveList(t *testing.T) {
	database.LoadEnv()

	list := List{
		Name: "Test list",
	}

	listID, publicCode, err := services.CreateList(&list)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	listID, err = services.RemoveList(listID)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	t.Log(listID, publicCode)
}

func TestUpdateList(t *testing.T) {
	database.LoadEnv()

	list := List{
		Name: "Test list",
	}

	listID, publicCode, err := services.CreateList(&list)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	list.Name = "Updated list"
	err = services.UpdateList(&list)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	listID, err = services.RemoveList(listID)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	t.Log(listID, publicCode)
}
