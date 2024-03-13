package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/config"
)

func main () {
	if err := config.Init(); err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	r.GET("/api/v1/planet", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	r.Run()
}