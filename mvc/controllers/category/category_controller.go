package categoryController

import (
	"mvc/dto"
	service "mvc/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//FUNCIONES CATEGORIA

func GetCategoryById(c *gin.Context) {
	log.Debug("Category id: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var categoryDto dto.CategoryDto
	categoryDto, err := service.CategoryService.GetCategoryById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, categoryDto)
}

func GetCategories(c *gin.Context) {
	var categoriesDto dto.CategoriesDto
	categoriesDto, err := service.CategoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, categoriesDto)

}
