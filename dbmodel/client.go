package dbmodel

// ClientData database model
type ClientData struct {
	Code      string  `db:"code" json:"code"`
	Name      string  `db:"name" json:"name"`
	IsActive  bool    `db:"is_active" json:"isActive"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdatedAt *string `db:"updated_at" json:"updated_at"`
}
