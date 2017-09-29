package ginmon

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)

// CounterHandler is a Gin middleware function that increments a
// global counter on each request.
func CounterHandler(ca *CounterAspect) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		fmt.Println("counter Handler")
		if(ctx.Keys[ctx.Request.URL.Path] != nil){
			ca.inc <- tuple{
				path: ctx.Keys[ctx.Request.URL.Path].(string),
				code: ctx.Writer.Status(),
			}
		}else{//handling invalid APIs
			if ctx.Keys == nil {
				ctx.Keys = make(map[string]interface{})
			}
			ctx.Keys[ctx.Request.URL.Path]= ctx.Request.URL.Path
			ca.inc <- tuple{
				path: ctx.Keys[ctx.Request.URL.Path].(string),
				code: ctx.Writer.Status(),
			}
		}

	}
}

type tuple struct {
	path string
	code int
}

// CounterAspect stores a counter
type CounterAspect struct {
	inc                  chan tuple
	internalRequestsSum  int
	internalRequests     map[string]int
	internalRequestCodes map[string]*RequestType
	RequestsSum          int            `json:"request_sum_per_minute"`
	Requests             map[string]int `json:"requests_per_minute"`
	RequestCodes map[string]*RequestType `json:"request_codes_per_api"`

}

type RequestType struct {
	RequestCodeCount         map[int]int  `json:"count_per_request_codes"`
	internalRequestCodesCount map[int]int
}

// NewCounterAspect returns a new initialized CounterAspect object.
func NewCounterAspect() *CounterAspect {
	ca := &CounterAspect{}
	ca.inc = make(chan tuple)
	ca.internalRequestsSum = 0
	ca.internalRequests = make(map[string]int, 0)
	ca.internalRequestCodes = make(map[string]*RequestType,0 )
	return ca
}

// StartTimer will call a forever loop in a goroutine to calculate
// metrics for measurements every d ticks. The parameter of this
// function should normally be 1 * time.Minute, if not it will expose
// unintuive JSON keys (requests_per_minute and
// request_sum_per_minute).
func (ca *CounterAspect) StartTimer(d time.Duration) {
	timer := time.Tick(d)
	go func() {
		for {
			select {
			case tup := <-ca.inc:
				fmt.Println("*********",tup.path)
				ca.increment(tup)
			case <-timer:
				//time.Sleep(5 * time.Second)
				ca.reset()

			}
		}
	}()
}

// GetStats to fulfill aspects.Aspect interface, it returns the data
// that will be served as JSON.
func (ca *CounterAspect) GetStats() interface{} {
	return *ca
}

// Name to fulfill aspects.Aspect interface, it will return the name
// of the JSON object that will be served.
func (ca *CounterAspect) Name() string {
	return "Counter"
}

// InRoot to fulfill aspects.Aspect interface, it will return where to
// put the JSON object into the monitoring endpoint.
func (ca *CounterAspect) InRoot() bool {
	return false
}

func (ca *CounterAspect) increment(tup tuple) {
	ca.internalRequestsSum++
	ca.internalRequests[tup.path]++
	if ca.internalRequestCodes[tup.path] == nil{
		ca.internalRequestCodes[tup.path] = &RequestType{internalRequestCodesCount : make(map[int]int,0 ),RequestCodeCount :make(map[int]int,0 ),
		}
		ca.internalRequestCodes[tup.path].internalRequestCodesCount[tup.code] = 0
	}
	ca.internalRequestCodes[tup.path].internalRequestCodesCount[tup.code]++
}

func (ca *CounterAspect) reset() {
	ca.RequestsSum = ca.internalRequestsSum
	ca.Requests = ca.internalRequests
	ca.RequestCodes = ca.internalRequestCodes
	for _, v := range ca.internalRequestCodes {
		v.RequestCodeCount = v.internalRequestCodesCount
	}
}
