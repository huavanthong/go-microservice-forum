package repositories

type OrderRepository struct {
	dbContext *OrderContext
}

func NewOrderRepository(dbContext *OrderContext) *OrderRepository {
	return &OrderRepository{dbContext: dbContext}
}

func (r *OrderRepository) GetOrdersByUserName(userName string) ([]Order, error) {
	var orderList []Order
	err := r.dbContext.Orders.Where("UserName = ?", userName).Find(&orderList).Error
	if err != nil {
		return nil, err
	}
	return orderList, nil
}
