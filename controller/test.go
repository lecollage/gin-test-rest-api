package controller

import (
	"fmt"
	"math/big"
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
	errBadParam := make(gin.H)
	errBadParam["error"] = "Bad param"

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

func calculateFactorial(number int64) string {
	var fact = new(big.Int)
	fact.MulRange(1, number)
	fmt.Println(fact)

	str := fact.String()
	sum := 0
	for _, val := range str {
		sum += int(val - '0')
	}

	return str
}

func CpuLoadSync(context *gin.Context) {
	errBadParam := make(gin.H)
	errBadParam["error"] = "Bad param"

	_, ok := context.Request.URL.Query()["number"]

	if !ok {
		context.JSON(http.StatusBadRequest, errBadParam)
		return
	}

	numberParam := context.Request.URL.Query()["number"][0]

	number, err := strconv.Atoi(numberParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, errBadParam)
		return
	}

	factorial := calculateFactorial(int64(number))

	fmt.Printf("work is done; factorial is %s\n", factorial)

	result := fmt.Sprintf("work is done; factorial is %s\n", factorial)

	context.JSON(http.StatusOK, gin.H{"result": result})
	return
}

func CpuLoadAsync(context *gin.Context) {
	errBadParam := make(gin.H)
	errBadParam["error"] = "Bad param"

	_, ok := context.Request.URL.Query()["ms"]

	if !ok {
		context.JSON(http.StatusBadRequest, errBadParam)
		return
	}

	msParam := context.Request.URL.Query()["ms"][0]

	durationMs, err := strconv.Atoi(msParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, errBadParam)
		return
	}

	stopChan := make(chan bool)

	go runCycle(stopChan)

	time.Sleep(time.Duration(durationMs) * time.Millisecond)
	stopChan <- true
	fmt.Println("Cycle has stopped.")

	context.JSON(http.StatusOK, gin.H{"result": "Work is done;"})
	return

}

func runCycle(stopChan <-chan bool) {
	for {
		select {
		case <-stopChan:
			return
		default:
			fmt.Println("Cycle is running...")
			time.Sleep(200 * time.Millisecond) // Sleep for 1 second before the next iteration.
		}
	}
}
