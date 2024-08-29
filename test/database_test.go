package test

import (
	"shopping-list-backend/pkg/database"
	"testing"
)

func TestConnection(t *testing.T) {
	err := database.LoadEnv()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	db, err := database.Connection()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	t.Logf("Connected to database: %v", db)

	defer db.Close()
}
