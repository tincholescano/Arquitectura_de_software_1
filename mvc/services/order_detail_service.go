package services

import (
	orderDetailClient "mvc/clients/order_detail"
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"

	log "github.com/sirupsen/logrus"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	InsertDetail(orderDetailDto dto.OrderDetailInsertDto, orderId int) (dto.OrderDetailResponseDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) InsertDetail(orderDetailDto dto.OrderDetailInsertDto, orderId int) (dto.OrderDetailResponseDto, e.ApiError) {

	var orderDetail model.OrderDetail
	orderDetail.OrderId = orderId
	orderDetail.ProductId = orderDetailDto.ProductId
	orderDetail.Quantity = orderDetailDto.Quantity
	orderDetail.CurrencyId = "ARS"
	orderDetail.Name = orderDetailDto.Name
	orderDetail.Price = orderDetailDto.Price

	orderDetail = orderDetailClient.InsertOrderDetail(orderDetail)

	var orderDetailResponseDto dto.OrderDetailResponseDto
	orderDetailResponseDto.OrderDetailId = orderDetail.OrderDetailId

	log.Debug(orderDetail)

	return orderDetailResponseDto, nil
}
