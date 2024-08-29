package models

type List struct {
	ID         *int   `json:"id"`
	PublicCode string `json:"public_code"`
	Name       string `json:"name"`
}

type Product struct {
	ListID    *int    `json:"list_id"`
	ID        *int    `json:"id"`
	Name      string  `json:"name"`
	Amount    *string `json:"amount"`
	Completed *bool   `json:"completed"`
	Notes     *string `json:"notes"`
}
