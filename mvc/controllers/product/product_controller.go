package productController

import (
	"mvc/dto"
	service "mvc/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//FUNCIONES PRODUCTOS

func GetProductById(c *gin.Context) {
	var productDto dto.ProductDto
	id, _ := strconv.Atoi(c.Param("product_id"))
	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {

	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByCategoryId(c *gin.Context) {

	var productsDto dto.ProductsDto
	id, _ := strconv.Atoi(c.Param("category_id"))
	productsDto, err := service.ProductService.GetProductsByCategoryId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productsDto)
}
