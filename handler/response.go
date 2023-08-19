package handler

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"msg":       msg,
		"errorCode": code,
	})
}
