package route

import (
	"net/http"
	"orm/handler"
	"orm/middleware"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	userHandler := handler.NewUserHandler()
	productHandler := handler.NewProductHandler()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "welcome to ecommerce")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")
	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}
	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)
		userProtectedRoutes.GET("/:id", userHandler.GetUser)
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)
	}

	productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/", productHandler.GetAllProduct)
		productRoutes.GET("/:id", productHandler.GetProduct)
		productRoutes.POST("/", productHandler.AddProduct)
		productRoutes.PUT("/:product", productHandler.UpdateProduct)
		productRoutes.DELETE("/:product", productHandler.DeleteProduct)
	}

	fileRoutes := r.Group("/file")
	{
		fileRoutes.POST("/single", handler.SingleFile)
		fileRoutes.POST("/multi", handler.MultipleFile)
	}
	return r.Run(address)
}
