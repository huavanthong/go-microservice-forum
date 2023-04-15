package queries

type GetOrdersListQueryHandler struct {
	orderRepository IOrderRepository
	mapper          IMapper
}

func NewGetOrdersListQueryHandler(orderRepository IOrderRepository, mapper IMapper) *GetOrdersListQueryHandler {
	return &GetOrdersListQueryHandler{
		orderRepository: orderRepository,
		mapper:          mapper,
	}
}

func (handler *GetOrdersListQueryHandler) Handle(request GetOrdersListQuery) ([]OrdersVm, error) {
	orderList, err := handler.orderRepository.GetOrdersByUserName(request.UserName)
	if err != nil {
		return nil, err
	}

	return handler.mapper.Map(orderList, []OrdersVm{}), nil
}
