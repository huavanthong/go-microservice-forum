package models

/************************ Define structure product ************************/
type Category struct {
	ID            int           `bson:"_id" json:"id" example:"001"`
	Name          string        `json:"name" bson:"name"`
	SubCategories []SubCategory `json:"subcategory" bson:"subcategory"`
	Description   string        `json:"description" bson:"description"`
	CreatedAt     string        `json:"created_at" bson:"created_at"`
	UpdatedAt     string        `json:"updated_at" bson:"updated_at"`
	DeleteAt      string        `json:"deleted_at" bson:"deleted_at"`
}

type SubCategory struct {
	CategoryType int    `bson:"categorytype" json:"categorytype" example:"001"`
	Name         string `json:"name" bson:"name"`
	Description  string `json:"description" bson:"description"`
	CreatedAt    string `json:"created_at" bson:"created_at"`
	UpdatedAt    string `json:"updated_at" bson:"updated_at"`
	DeleteAt     string `json:"deleted_at" bson:"deleted_at"`
}
