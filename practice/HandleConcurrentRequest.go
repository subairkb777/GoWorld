package main

import(
	"fmt"
	"time"
	"sync"
	"math/rand"

)

var r *rand.Rand // Rand for this package.

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func getRandomString()string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := ""
	for i := 0; i < 5; i++ {
		index := r.Intn(len(chars))
		result += chars[index: index+1]
	}
	return result
}


type Job struct {
	processingTime       int
	str string
}

type Result struct {
	job         Job
	asciiSum int
}



var resultSum int
var jobs = make(chan Job, 1000)
var results = make(chan Result, 1000)

func findCharSum(number string, sleepTime int) int {
	sum := 0
	//find ASCII Sum
	// sleep process for sleepTime
	return sum
}

//creates a worker Goroutine.
func worker(wg *sync.WaitGroup) {

	// read the job queued in jobs channel
	for job := range jobs {
		// for each job in "jobs channel" calculate sum and store it in results channel
		output := Result{job, findCharSum(job.str, job.processingTime)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	//close(results)
}

//Allocate job to workers in pool
func allocate( ) {
	for {
		allocateJob()
		time.Sleep(1 * time.Minute)
	}

}

func allocateJob(){
	for i := 0; i < 1000; i++ {
		job := Job{i, getRandomString()}
		jobs <- job
	}
}

// read Result from channel
func result() {
	for result := range results {
		//do :read results
		//store sum in result to global variable resultSum
	}
}

func main(){

	go allocate()
	//go result()
	//noOfWorkers := 100
	//createWorkerPool(noOfWorkers)
}
