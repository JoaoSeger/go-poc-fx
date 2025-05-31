package presentation

import (
	"net/http"
	"strconv"

	"go-poc-fx/internal/user/domain"

	"github.com/gin-gonic/gin"
)

// UserController handle HTTP requests for user operations
type UserController struct {
	userService domain.UserService
}

// NewUserController creates a new UserController instance
func NewUserController(userService domain.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// RegisterRoutes registers all user routes
func (uc *UserController) RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", uc.CreateUser)
		userGroup.GET("/:id", uc.GetUser)
		userGroup.GET("", uc.GetAllUsers)
	}
}

// CreateUser handles POST /users
func (uc *UserController) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.CreateUser(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUser handles GET /users/:id
func (uc *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := uc.userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers handles GET /users
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
