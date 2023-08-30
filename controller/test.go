package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello, world :)")
}

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, "Pong")
}

func Sleep(context *gin.Context) {
	errBadParam := struct{ error string }{error: "Bad param"}

	_, ok := context.Request.URL.Query()["s"]

	if !ok {
		context.JSON(http.StatusBadRequest, errBadParam)
		return
	}

	sParam := context.Request.URL.Query()["s"][0]

	s, err := strconv.Atoi(sParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, errBadParam)
		return
	}

	duration := time.Duration(s) * time.Second

	time.Sleep(duration)

	context.JSON(http.StatusOK, gin.H{"result": "Work is done"})
	return
}
