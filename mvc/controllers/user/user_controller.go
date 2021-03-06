package userController

import (
	"mvc/dto"
	service "mvc/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//USUARIO

func GetUserById(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))
	var userDto dto.UserDto
	id, _ := strconv.Atoi(c.Param("id"))
	userDto, err := service.UserService.GetUserById(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {

	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UserInsert(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	log.Debug(userDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.UserService.InsertUser(userDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	c.BindJSON(&loginDto)
	log.Debug(loginDto)

	var loginResponseDto dto.LoginResponseDto
	loginResponseDto, err := service.UserService.Login(loginDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, loginResponseDto)
}
