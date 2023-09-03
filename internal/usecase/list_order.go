package usecase

import (
	"github.com/bvaledev/go-clean-arch/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var ordersList []OrderOutputDTO

	for _, item := range orders {
		item.CalculateFinalPrice()
		order := OrderOutputDTO{
			ID:         item.ID,
			Price:      item.Price,
			Tax:        item.Tax,
			FinalPrice: item.FinalPrice,
		}
		ordersList = append(ordersList, order)
	}

	return ordersList, nil
}
