package main


import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-gomonitor"
	"github.com/szuecs/gin-gomonitor/aspects"
	"gopkg.in/mcuadros/go-monitor.v1/aspects"
	"time"
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

	router.GET("/API1",Validator("/API1"), func(ctx *gin.Context) {
		//time.Sleep(10 *time.Second)
		ctx.JSON(http.StatusOK, gin.H{ "hello": "world API1"})
	})

	router.GET("/API2",Validator("/API2"), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "world API2"})
	})

	router.GET("/API3",Validator("/API3"), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "world API3"})
	})

	router.GET("/API5",Validator("/API5"), func(ctx *gin.Context) {
		param := ctx.DefaultQuery("param", "Guest")
		if(param == "1"){
			ctx.JSON(http.StatusOK, gin.H{ "hello": "world API5"})
		}else {
			ctx.JSON(http.StatusForbidden, gin.H{ "hello": "Forbidden"})
		}

	})

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/API6/:name",Validator("/API6/:name"), func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	log.Fatal(router.Run(":8080"))
}

func Validator(path interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys[c.Request.URL.Path] = path
	}
}