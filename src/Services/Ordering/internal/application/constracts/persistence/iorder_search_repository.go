package constracts_persistence

type IOrderRepository interface {
	IAsyncRepository
	GetOrdersByUserName(userName string) ([]Order, error)
}
