package services

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"shopping-list-backend/internal/models"
	"shopping-list-backend/pkg/database"

	_ "github.com/lib/pq"
)

func generateRandomCode(length int, db *sql.DB) (string, error) {
	for {
		// Generate a random code
		b := make([]byte, length)
		_, err := rand.Read(b)
		if err != nil {
			return "", err
		}
		code := hex.EncodeToString(b)

		var exists bool
		err = db.QueryRow(`
            SELECT EXISTS (
                SELECT 1 FROM lists WHERE public_code = $1
            )
        `, code).Scan(&exists)
		if err != nil {
			return "", err
		}

		if !exists {
			return code, nil
		}
	}
}

func CreateList(list *models.List) (int, string, error) {
	var lastInsertID int

	db, err := database.Connection()
	if err != nil {
		return -1, "", err
	}

	randomCode, err := generateRandomCode(8, db)
	if err != nil {
		return -1, "", err
	}

	err = db.QueryRow("INSERT INTO lists (public_code, name) VALUES ($1, $2) RETURNING id", randomCode, list.Name).Scan(&lastInsertID)
	if err != nil {
		return -1, "", err
	}

	return lastInsertID, randomCode, nil
}

func RemoveList(listID int) (int, error) {
	db, err := database.Connection()
	if err != nil {
		return -1, err
	}

	_, err = db.Exec("DELETE FROM lists WHERE id = ($1)", listID)
	if err != nil {
		return -1, err
	}

	return listID, nil
}

func UpdateList(list *models.List) error {
	db, err := database.Connection()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE lists SET name = $1 WHERE public_code = $2", list.Name, list.PublicCode)
	if err != nil {
		return err
	}

	return nil
}

func GetListByCode(publicCode string) (*models.List, error) {
	fmt.Printf("Getting list with code: %s\n", publicCode)
	db, err := database.Connection()
	if err != nil {
		return nil, err
	}

	list := models.List{}

	err = db.QueryRow("SELECT id, public_code, name FROM lists WHERE public_code = $1", publicCode).Scan(&list.ID, &list.PublicCode, &list.Name)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func GetFirstList() (int, error) {
	db, err := database.Connection()
	if err != nil {
		return -1, err
	}

	var id int

	err = db.QueryRow("SELECT id FROM lists LIMIT 1").Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}
