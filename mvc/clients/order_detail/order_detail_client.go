package orderDetail

import (
	"mvc/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

//FUNCIONES DETALLE DE ORDENES

func GetOrderDetailById(id int) model.OrderDetail {
	var orderDetail model.OrderDetail
	Db.Where("order_detail_id = ?", id).First(&orderDetail)
	log.Debug("OrderDetail: ", orderDetail)

	return orderDetail
}

func InsertOrderDetail(orderDetail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&orderDetail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("OrderDetail Created: ", orderDetail.OrderDetailId)
	return orderDetail
}
