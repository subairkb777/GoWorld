package main

import (
	"time"
)

func main() {

	//Handle Counter
	for {
		countHandler("http://localhost:9000/Counter", "678053105476", "FinalCounterMatrics")
		//Handle requestTime
		requestTimeHandler("http://localhost:9000/RequestTime", "678053105476", "FinalRequestTimeMatrics")

		time.Sleep(10 * time.Second)
	}


}


