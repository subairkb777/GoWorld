package main


import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-gomonitor"
	"github.com/szuecs/gin-gomonitor/aspects"
	"gopkg.in/mcuadros/go-monitor.v1/aspects"
)

func main1() {
	// initialize CounterAspect and reset every minute
	counterAspect2 := ginmon.NewCounterAspect()
	counterAspect2.StartTimer(5 * time.Second)
	asps2 := []aspects.Aspect{counterAspect2}
	router := gin.New()
	// register CounterAspect middleware
	// test: curl http://localhost:9000/Counter
	router.Use(ginmon.CounterHandler(counterAspect2))

	// start metrics endpoint
	gomonitor.Start(9000, asps2)
	// last middleware
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "world"})
	})

	log.Fatal(router.Run(":8080"))
}