package main

import (
	"github.com/szuecs/gin-gomonitor/aspects"
	"time"
	"gopkg.in/mcuadros/go-monitor.v1/aspects"
	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-gomonitor"
	"net/http"
	"log"
)

func main4() {
	// initialize RequestTimeAspect and calculate every 5 seconds
	requestAspect1 := ginmon.NewRequestTimeAspect()
	requestAspect1.StartTimer(5 * time.Second)
	asps1 := []aspects.Aspect{requestAspect1}

	router := gin.New()
	// register RequestTimeAspect middleware
	// test: curl http://localhost:9000/RequestTime
	router.Use(ginmon.RequestTimeHandler(requestAspect1))
	// start metrics endpoint
	gomonitor.Start(9000, asps1)
	// last middleware
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "world"})
	})

	log.Fatal(router.Run(":8080"))
}