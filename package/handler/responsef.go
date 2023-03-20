package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
}
