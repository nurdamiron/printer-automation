package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/nurdamiron/printer-automation/internal/service"
)

type UserHandler struct {
    userService service.UserService
}

func NewUserHandler(u service.UserService) *UserHandler {
    return &UserHandler{userService: u}
}
//
func (h *UserHandler) CreateUser(c *gin.Context) {
    type reqBody struct {
        Username string `json:"username" binding:"required"`
    }
    var req reqBody
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.userService.CreateUser(c.Request.Context(), req.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := h.userService.GetUser(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}
