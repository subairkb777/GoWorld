package main

/*import (
	health "github.com/gocraft/health"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-health"
	//"github.com/gocraft/health/sinks/bugsnag"
	"log"
	"time"
)
var stream = health.NewStream()

// In your main func, initiailze the stream with your sinks.
func main() {
	// Log to stdout! (can also use WriterSink to write to a log file, Syslog, etc)

	// setup stream with sinks
	//router := gin.New()

	startTime := time.Now()

	router := gin.Default()

	hstream := ghealth.NewStream("", "127.0.0.1:5020")

	router.Use(ghealth.Health(hstream,false))

	var jobRoot *health.Job
	root :=func(ctx *gin.Context) {
		jobRoot = hstream.NewJob("test_post")
		ctx.JSON(http.StatusOK, gin.H{ "hello": "root1"})
	}
	/*test := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "test1"})
	}*/


	//job := hstream.NewJob("test_post")
	//job.Event()
	//Send errors to bugsnag!
	//stream.AddSink(bugsnag.NewSink(&bugsnag.Config{APIKey: "myApiKey"}))

	/*if err := router.GET("/", root); err == nil {
		jobRoot.Timing("test_post", time.Since(startTime).Nanoseconds())
		jobRoot.Complete(health.Success)
		log.Fatal("main function ",err)
	}
	/*if err := router.GET("/test", test); err == nil {
		job.Complete(health.Success)
		log.Fatal("test function ",err)
	}*/

	//router.Run(":8080")

	// Now that your stream is setup, start a web server or something...
//}
