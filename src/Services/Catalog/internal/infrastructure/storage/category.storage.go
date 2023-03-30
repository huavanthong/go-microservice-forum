package storage

// CategoryRepository is an interface for managing product based as CRUD operation
type CategoryRepository interface {
	Create(userName string) (*entities.ShoppingCart, error)
	GetByUserName(userName string) (*entities.ShoppingCart, error)
	Update(basket *entities.ShoppingCart) (*entities.ShoppingCart, error)
	Delete(userName string) error
}
