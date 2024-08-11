package user

import (
    "my-hexagonal-api/internal/domain/user"
    "github.com/gin-gonic/gin"
    "net/http"
)

type UserController struct {
    service *user.UserService
}

func NewUserController(service *user.UserService) *UserController {
    return &UserController{service: service}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
    var req struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newUser, err := uc.service.RegisterUser(req.Name, req.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, newUser)
}
