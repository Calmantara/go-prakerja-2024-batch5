package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Middleware1(ctx *gin.Context) {
	fmt.Println("middleware 1")
	ctx.Next()
}

func Middleware2(ctx *gin.Context) {
	fmt.Println("middleware 2")
	ctx.Next()
}

func Middleware3(ctx *gin.Context) {
	fmt.Println("middleware 3")
	ctx.Next()
}
