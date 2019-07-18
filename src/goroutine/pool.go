package main

import (
	"time"
	"sync"
	"math/rand"
	"fmt"
)

type Job struct {
	id			int
	randomno	int
}

type Result struct {
	job			Job
	sumofdigits int
}

// 接收作业和写入结果的缓冲信道
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// digits 函数的任务实际上就是计算整数的每一位之和，最后返回该结果。为了模拟出 digits 在计算过程中花费了一段时间，我们在函数内添加了两秒的休眠时间
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(time.Millisecond * 100)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// 创建了一个 Go 协程的工作池
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		//println("----+",i)
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// 作业分配给工作者
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 读取 results 信道和打印输出的函数
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}


//- 创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
//- 将作业添加到该输入型缓冲信道中。
//- 作业完成后，再将结果写入一个输出型缓冲信道。
//- 从输出型缓冲信道读取并打印结果。

func main()  {
	startTime := time.Now()
	noOfJobs := 1000
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

}