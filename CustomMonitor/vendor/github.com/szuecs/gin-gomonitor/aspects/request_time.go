package ginmon

import (
	"time"

	"github.com/gin-gonic/gin"
)

type APIPath struct {
	path string
	requestTime time.Time
}

// RequestTimeAspect, exported fields are used to store json
// fields. All fields are measured in nanoseconds.
type RequestTimeAspect struct {
	inc                  chan APIPath
	counter					int
	lastMinuteRequestTimes []float64
	Count                  int       `json:"count"`/*
	Min                    float64   `json:"min"`
	Max                    float64   `json:"max"`
	Mean                   float64   `json:"mean"`
	Stdev                  float64   `json:"stdev"`
	P90                    float64   `json:"p90"`
	P95                    float64   `json:"p95"`
	P99                    float64   `json:"p99"`*/
	Timestamp              time.Time `json:"timestamp"`
	RequestAPITime map[string][]float64
	AvgTime    map[string]float64
}

// NewRequestTimeAspect returns a new initialized RequestTimeAspect
// object.
func NewRequestTimeAspect() *RequestTimeAspect {
	rt := &RequestTimeAspect{}
	rt.inc = make(chan APIPath)
	rt.lastMinuteRequestTimes = make([]float64, 0)
	rt.Timestamp = time.Now()
	rt.RequestAPITime = make(map[string][]float64, 0)
	rt.AvgTime = make(map[string] float64, 0)
	return rt
}

// StartTimer will call a forever loop in a goroutine to calculate
// metrics for measurements every d ticks.
func (rt *RequestTimeAspect) StartTimer(d time.Duration) {
	timer := time.Tick(d)
	go func() {
		for {
			select{
				case tup := <-rt.inc:
					took := time.Now().Sub(tup.requestTime)
					rt.add(float64(took),tup)
				case <-timer:
					rt.calculate()
			}
		}
	}()
}


// GetStats to fulfill aspects.Aspect interface, it returns the data
// that will be served as JSON.
func (rt *RequestTimeAspect) GetStats() interface{} {
	return rt
}

// Name to fulfill aspects.Aspect interface, it will return the name
// of the JSON object that will be served.
func (rt *RequestTimeAspect) Name() string {
	return "RequestTime"
}

// InRoot to fulfill aspects.Aspect interface, it will return where to
// put the JSON object into the monitoring endpoint.
func (rt *RequestTimeAspect) InRoot() bool {
	return false
}

// RequestTimeHandler is a middleware function to use in Gin
func RequestTimeHandler(rt *RequestTimeAspect) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()

		if(c.Keys[c.Request.URL.Path] != nil){
			rt.inc <- APIPath{
				path: c.Keys[c.Request.URL.Path].(string),
				requestTime: now,
			}
		}else{//handling invalid APIs
			if c.Keys == nil {
				c.Keys = make(map[string]interface{})
			}
			c.Keys[c.Request.URL.Path]= c.Request.URL.Path
			rt.inc <- APIPath{
				path: c.Keys[c.Request.URL.Path].(string),
				requestTime: now,
			}
		}


	}
}



func (rt *RequestTimeAspect) add(n float64,tup APIPath) {

	if rt.RequestAPITime[tup.path] == nil{
		rt.RequestAPITime[tup.path] = make([]float64, 0)
	}
	rt.RequestAPITime[tup.path] = append(rt.RequestAPITime[tup.path], n)

}

func (rt *RequestTimeAspect) calculate() {

	for k, v := range rt.RequestAPITime {
		l := len(v)
		avg := mean(v, l)
		rt.AvgTime[k]=avg
	}
	rt.RequestAPITime = make(map[string][]float64, 0)

}

