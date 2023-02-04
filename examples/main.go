package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/PaulShpilsher/concurrent-go/runner"
)

func worker(id int) {
	fmt.Println("worker start", id)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Println("worker end  ", id)
}

func main() {
	var limit int
	var tasks int

	flag.IntVar(&limit, "limit", 2, "maxiumum number of goroutines allowed to run concurrently")
	flag.IntVar(&tasks, "tasks", 5, "the number of tasks we need to execute")
	flag.Parse()

	fmt.Printf("Executing %d tasks with concurrency limit %d\n", tasks, limit)

	rand.Seed(time.Now().UnixNano())
	r, _ := runner.NewConcurrencyRunner(limit)
	for i := 0; i < tasks; i++ {
		id := i
		r.Run(func() { worker(id) })
	}

	r.Close() // waits for all pending tasks to complete and closes runner
}
