package services

import (
	orderClient "mvc/clients/order"
	productClient "mvc/clients/product"
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"
	"time"

	log "github.com/sirupsen/logrus"
)

type orderService struct{}

type orderServiceInterface interface {
	InsertOrder(orderDto dto.OrderInsertDto) (dto.OrderResponseDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) InsertOrder(orderInsertDto dto.OrderInsertDto) (dto.OrderResponseDto, e.ApiError) {

	var order model.Order
	var total float32
	var orderResponseDto dto.OrderResponseDto
	total = 0
	order.UserId = orderInsertDto.UserId
	order.Date = time.Now().Format("2006.01.02 15:04:05")
	for i := 0; i < len(orderInsertDto.OrderDetails); i++ {
		var product model.Product
		detail := orderInsertDto.OrderDetails[i]
		product = productClient.GetProductById(detail.ProductId)
		if product.Stock < detail.Quantity {
			orderResponseDto.OrderId = -1
			return orderResponseDto, e.NewConflictApiError("Not enough stock on product: " + product.Name)
		}

		total += (detail.Price * float32(detail.Quantity))

	}

	for i := 0; i < len(orderInsertDto.OrderDetails); i++ {
		detail := orderInsertDto.OrderDetails[i]
		productClient.RemoveStock(detail.ProductId, detail.Quantity)
	}

	order.Total = total
	order.CurrencyId = "ARS"

	order = orderClient.InsertOrder(order)

	orderResponseDto.OrderId = order.OrderId

	log.Debug(order)
	return orderResponseDto, nil
}
