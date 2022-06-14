package app

import (
	categoryController "mvc/controllers/category"
	orderController "mvc/controllers/order"
	orderDetailController "mvc/controllers/order_detail"
	productController "mvc/controllers/product"
	userController "mvc/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Login
	router.POST("/login", userController.Login)

	// Usuarios
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

	// Productos
	router.GET("/product/:product_id", productController.GetProductById)
	router.GET("/products", productController.GetProducts)
	router.GET("/products/:category_id", productController.GetProductsByCategoryId)

	// Ordenes
	router.GET("/order/:id", orderController.GetOrderById)
	router.POST("/order", orderController.OrderInsert)

	// Detalle Ordenes
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.POST("/orderDetail", orderDetailController.OrderDetailInsert)

	// Categorias
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/categories", categoryController.GetCategories)

	log.Info("Finishing mappings configurations")
}
