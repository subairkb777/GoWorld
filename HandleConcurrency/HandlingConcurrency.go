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

func generateRandom() int{
	rand.Seed(time.Now().Unix())
	return rand.Intn(600) + 300
}

type Job struct {
	processingTime       int
	str string
}

type Result struct {
	job         Job
	asciiSum int
}


var mutex = &sync.Mutex{}
var resultSum int


func findCharSum(str string, sleepTime int) int {
	sum := 0
	for i:=0;i<len(str);i++{
		sum += int(str[0])
	}
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	return sum
}

//creates a worker Goroutine.
func worker(jobs chan Job,results chan Result) {

	// read the job queued in jobs channel
	for job := range jobs {
		// for each job in "jobs channel" calculate sum and store it in results channel
		output := Result{job, findCharSum(job.str, job.processingTime)}
		results <- output
	}
}

func createWorkerPool(noOfWorkers int,jobs chan Job,results chan Result) {
	for i := 0; i < noOfWorkers; i++ {
		go worker(jobs,results)
	}
}


func allocateJob(maxJobAtATime int,jobs chan Job){
	for i := 0; i < maxJobAtATime; i++ {
		job := Job{(generateRandom()+i)%1000, getRandomString()}
		jobs <- job
	}
}

// read Result from channel done chan bool
func result(results chan Result) {
	for result := range results {
		mutex.Lock()
		resultSum += result.asciiSum
		mutex.Unlock()

	}
}

func printSum(){
	fmt.Println("Sum = ", resultSum)
	resultSum =0
}

func startBatch(newBatch chan bool,jobs chan Job,results chan Result){
	fmt.Println("Allocation of new job")
	maxJobAtATime := 1000 //1000 request
	go allocateJob(maxJobAtATime,jobs)
	go result(results)
	noOfWorkers := 100
	createWorkerPool(noOfWorkers,jobs,results)
	newBatch <- true

}
func main(){

	numberOfbatch := 3
	for i:=0;i<numberOfbatch;i++{
		fmt.Println("Batch Number ",i+1)
		var jobs = make(chan Job, 1000)
		var results = make(chan Result, 1000)
		var newBatch = make(chan bool)
		go startBatch(newBatch,jobs,results)
		<- newBatch
		time.Sleep(10 * time.Second)
		printSum()
	}

}
