package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// CounterAspect stores a counter
type CounterAspect struct {
	counter string `json:"Counter"`
	RequestsSum          int            `json:"request_sum_per_minute"`
	Requests             map[string]int `json:"requests_per_minute"`
	RequestCodes map[string]*RequestType `json:"request_codes_per_api"`

}

type RequestType struct {
	RequestCodeCount         map[int]int  `json:"count_per_request_codes"`
	internalRequestCodesCount map[int]int
}


func main() {

	res, err := http.Get("http://localhost:9000/Counter")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Body")
	fmt.Println(string(body[:]))
	s, err := getStations([]byte(body))

	fmt.Println(s)


}
func getStations(body []byte) (*CounterAspect, error) {
	var s CounterAspect
	err := json.Unmarshal(body, &s)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return &s, err
}