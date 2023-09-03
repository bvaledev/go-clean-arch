package service

import (
	"context"

	"github.com/bvaledev/go-clean-arch/internal/infra/grpc/pb"
	"github.com/bvaledev/go-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (c *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderListResponse, error) {
	outputResult, err := c.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var output []*pb.CreateOrderResponse
	for _, item := range outputResult {
		order := &pb.CreateOrderResponse{
			Id:         item.ID,
			Price:      float32(item.Price),
			Tax:        float32(item.Tax),
			FinalPrice: float32(item.FinalPrice),
		}
		output = append(output, order)
	}
	return &pb.OrderListResponse{Orders: output}, nil
}
