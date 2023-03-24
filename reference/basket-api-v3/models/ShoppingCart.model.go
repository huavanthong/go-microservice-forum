package models

/************************ Define structure product ************************/
type ShoppingCart struct {
	UserName   string `bson:"username" json:"username" example:"username"`
	Items      []ShoppingCartItem
	TotalPrice float32
}

func (sc *ShoppingCart) getToTalPrice() error {

}
