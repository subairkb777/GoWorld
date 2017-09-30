package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func openAPI(url string){
	response, err := http.Get(url)

	if err != nil {
		fmt.Print("error in opening 1 :",err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error in response 1",err)
	}
	fmt.Println(string(responseData))

}

func main() {
	for {
		go openAPI("http://127.0.0.1:8080/API1")
		//time.Sleep(1 * time.Second)
		go openAPI("http://127.0.0.1:8080/API2")
		//time.Sleep(1 * time.Second)
		go openAPI("http://127.0.0.1:8080/API3")
		//time.Sleep(1 * time.Second)
		go openAPI("http://127.0.0.1:8080/API4")
		//time.Sleep(1 * time.Second)
		go openAPI("http://127.0.0.1:8080/API5?param=1")
		//time.Sleep(1 * time.Second)
		go openAPI("http://127.0.0.1:8080/API5?param=2")
		go openAPI("http://127.0.0.1:8080/API5?param=2")
		//time.Sleep(1 * time.Second)
		go openAPI("http://127.0.0.1:8080/API6/subair")
		//time.Sleep(1 * time.Second)
		go openAPI("http://127.0.0.1:8080/API6/john")
		time.Sleep(1* time.Second)
	}




}