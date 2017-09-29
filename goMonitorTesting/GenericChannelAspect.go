package main

import (
	"github.com/szuecs/gin-gomonitor/aspects"
	"time"
	"gopkg.in/mcuadros/go-monitor.v1/aspects"
	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-gomonitor"
)

func main9() {
	// initialize GenericChannelAspect and calculate every 3 seconds
	genericAspect := ginmon.NewGenericChannelAspect("generic")
	genericAspect.StartTimer(3 * time.Second)
	genericCH := genericAspect.SetupGenericChannelAspect()
	asps := []aspects.Aspect{genericAspect}

	router := gin.New()
	// register GenericChannelAspect middleware
	// test: curl http://localhost:9000/generic
	// start metrics endpoint
	gomonitor.Start(9000, asps)
	// catch panics as last middleware
	router.Use(gin.Recovery())

	// send a lot of data concurrently to the monitoring data channel
	i := 0
	go func() {
		for {
			i++
			genericCH <- ginmon.DataChannel{Name: "foo", Value: float64(i)}
		}
	}()
	j := 0
	go func() {
		for {
			j++
			genericCH <- ginmon.DataChannel{Name: "bar", Value: float64(j % 5)}
		}
	}()

	router.Run(":8080")
}
