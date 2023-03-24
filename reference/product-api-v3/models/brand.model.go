package models

type Brand struct {
	ID        int
	Name      string
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	DeleteAt  string `json:"deleted_at" bson:"deleted_at"`
}
