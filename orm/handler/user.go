package handler

import (
	"fmt"
	"net/http"
	"orm/model"
	"orm/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
	GetAllUser(*gin.Context)
	SignInUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() UserHandler {
	return &userHandler{
		repo: repository.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytesPass := []byte(*pass)
	hpass, _ := bcrypt.GenerateFromPassword(bytesPass, bcrypt.DefaultCost)
	*pass = string(hpass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

// GetAllUser implements UserHandler
func (h *userHandler) GetAllUser(ctx *gin.Context) {
	fmt.Println(ctx.Get("userID"))
	user, err := h.repo.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// GetUser implements UserHandler
func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.repo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// SignInUser implements UserHandler
func (h *userHandler) SignInUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := h.repo.GetByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "No user found"})
		return
	}

	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := GenerateToken(dbUser.ID)

		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully signedIn", "token": token})
		return

	}
	ctx.JSON(http.StatusInternalServerError,
		gin.H{"msg": "Password not matched"})
	return

}

// AddUser implements UserHandler
func (h *userHandler) AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	hashPassword(&user.Password)
	user, err := h.repo.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

// UpdateUser implements UserHandler
func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(intID)
	user, err = h.repo.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// DeleteUser implements UserHandler
func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var user model.User
	id := ctx.Param("user")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)
	user, err := h.repo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
