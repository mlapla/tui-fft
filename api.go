package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	math "math"
	rand "math/rand"
)

const SERVER_PORT = 8000

type Signal struct {
	Re []float64
	Im []float64
}

func makeSource(t int) []complex128 {
	n := 400
	data := make([]complex128, n)
	
	for i := range data {
		data[i] = complex(
			math.Sin(float64(2*i-t)/5) + math.Pow(rand.Float64(),2),
			math.Sin(float64(i-t+ 1)/5) + math.Pow(rand.Float64(),2))
	}

	return data
}


func getSignal(ctx *gin.Context) {

	sig := makeSource(time.Now().Second())

	flat_re, flat_im := Flatten(sig)

	model := &Signal { flat_re, flat_im }

	ctx.IndentedJSON(http.StatusOK, model)
}


func RunApi(logger* log.Logger) {

	router := gin.New()
	
	log.Println("Defining routes")
	router.GET("/signals/1", getSignal)

	log.Println("Booting Server...")
	url := fmt.Sprintf("localhost:%d", SERVER_PORT)
	router.Run(url)
}
