package handler

import (
	"fmt"
	"net/http"
	"orm/model"
	"orm/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(*gin.Context)
	GetProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

// GetAllProduct implements ProductHandler
func (h *productHandler) GetAllProduct(ctx *gin.Context) {
	fmt.Println(ctx.Get("productID"))
	fmt.Println("GetAllProduct")
	product, err := h.repo.GetAllProduct()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// GetProduct implements ProductHandler

func (h *productHandler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println("GetProduct")
	fmt.Println("id", id)
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := h.repo.GetProduct(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// AddProduct implements ProductHandler
func (h *productHandler) AddProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	product, err := h.repo.AddProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// UpdateProduct implements ProductHandler
func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("product")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(intID)
	product, err = h.repo.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct implements ProductHandler
func (h *productHandler) DeleteProduct(ctx *gin.Context) {
	var product model.Product
	id := ctx.Param("product")
	intID, _ := strconv.Atoi(id)
	product.ID = uint(intID)
	product, err := h.repo.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
