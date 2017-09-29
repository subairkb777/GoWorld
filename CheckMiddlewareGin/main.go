package main

import (
	"log"
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-gomonitor"
	"github.com/szuecs/gin-gomonitor/aspects"
	"gopkg.in/mcuadros/go-monitor.v1/aspects"
	"time"
	"fmt"
)

func main() {
	router := gin.New()

	// initialize CounterAspect and reset every minute
	counterAspect := ginmon.NewCounterAspect()
	counterAspect.StartTimer(5 * time.Second)

	requestAspect := ginmon.NewRequestTimeAspect()
	requestAspect.StartTimer(5 * time.Second)

	asps := []aspects.Aspect{counterAspect,requestAspect}

	router.Use(ginmon.CounterHandler(counterAspect))

	router.Use(ginmon.RequestTimeHandler(requestAspect))

	// start metrics endpoint
	gomonitor.Start(9000, asps)
	// last middleware
	router.Use(gin.Recovery())

	//"/API1/:name/:hello"
	router.GET("/API1/hello:name/:hello",Validator("/API1/hello:name/:hello"), fun1,fun2,fun3,fun4,func(ctx *gin.Context) {
		name := ctx.Param("name")
		hello := ctx.Param("hello")
		fmt.Println("name:",name," hello:",hello)
		ctx.JSON(http.StatusOK, gin.H{ hello: name})
	})



	log.Fatal(router.Run(":8080"))
}
func fun1(c *gin.Context){

	fmt.Println("middleware 1",c.Request.URL.Path)
	for k, v := range c.Keys {
		fmt.Println("Key :",k," Value :",v)
	}
}
func fun2(c *gin.Context){
	fmt.Println("middleware 2")
}
func fun3(c *gin.Context){
	fmt.Println("middleware 3")
}
func fun4(c *gin.Context){
	fmt.Println("middleware 4")
}


func Validator(path interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("URL 1",path)
		fmt.Println("Path 1",c.Keys[c.Request.URL.Path])
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys[c.Request.URL.Path] = path
		fmt.Println("URL 1",c.Keys[c.Request.URL.Path])
	}
}