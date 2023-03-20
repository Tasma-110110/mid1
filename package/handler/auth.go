package handler

import (
	"github.com/Tasma-110110/mid1-prj"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input mid1.User
	if err := c.BindJSON(&input); err != nil {
		return
	}

}

type signInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {

}
