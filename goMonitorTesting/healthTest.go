package mainpackage

import (
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-health"
	"github.com/gocraft/health/sinks/bugsnag"
)
main

import (
health "github.com/gocraft/health"
"net/http"
"os"
"github.com/gin-gonic/gin"
//"time"
"github.com/utrack/gin-health"
"time"
"github.com/gocraft/health/sinks/bugsnag"
)
var stream = health.NewStream()
/*func main() {

	// setup stream with sinks
	//router := gin.New()

	router := gin.Default()

	stream.AddSink(&health.WriterSink{os.Stdout})

	root :=func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "root1"})
	}
	test := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "test1"})
	}
	// First, you need to create a new stream...

	// Simplest sink, stdout only
	//hstream := ghealth.NewStream("", "", "")

	// STDOUT and JSON sinks; creates independent http server on port 5020
	//hstream := ghealth.NewStream("", "", "127.0.0.1:5020")

	// StatsD and JSON sinks
	hstream := ghealth.NewStream("", "", "127.0.0.1:5020")

	// It's a standard *health.Stream, so you can do anything you want!
	hstream.AddSink(&health.WriterSink{os.Stdout})


	router.Use(ghealth.Health(hstream,false))

	router.GET("/", root)
	router.GET("/test", test )
	router.Run(":8080")
}*/

// In your main func, initiailze the stream with your sinks.
func main() {
	// Log to stdout! (can also use WriterSink to write to a log file, Syslog, etc)
	stream.AddSink(&health.WriterSink{os.Stdout})


	// Expose instrumentation in this app on a JSON endpoint that healthd can poll!
	sink := health.NewJsonPollingSink(time.Minute, time.Minute*5)
	stream.AddSink(sink)

	// setup stream with sinks
	//router := gin.New()

	router := gin.Default()

	stream.AddSink(&health.WriterSink{os.Stdout})

	root :=func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "root1"})
	}
	test := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "hello": "test1"})
	}
	// First, you need to create a new stream...

	// Simplest sink, stdout only
	//hstream := ghealth.NewStream("", "", "")

	// STDOUT and JSON sinks; creates independent http server on port 5020
	//hstream := ghealth.NewStream("", "", "127.0.0.1:5020")

	// It's a standard *health.Stream, so you can do anything you want!
	//hstream.AddSink(&health.WriterSink{os.Stdout})


	router.Use(ghealth.Health(stream,false))

	sink.StartServer(":9000")

	// Send errors to bugsnag!
	stream.AddSink(bugsnag.NewSink(&bugsnag.Config{APIKey: "myApiKey"}))

	router.GET("/", root)
	router.GET("/test", test )
	router.Run(":8080")

	// Now that your stream is setup, start a web server or something...
}
