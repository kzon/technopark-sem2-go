package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}

	var in, out1, out2 chan interface{}
	in = make(chan interface{}, MaxInputDataLen)

	for i := 0; i < len(jobs); i += 2 {
		out1 = make(chan interface{}, MaxInputDataLen)
		out2 = make(chan interface{}, MaxInputDataLen)

		wg.Add(1)
		go runJob(jobs[i], in, out1, wg)
		if i+1 < len(jobs) {
			wg.Add(1)
			go runJob(jobs[i+1], out1, out2, wg)
		}

		in = out2
	}

	wg.Wait()
}

func runJob(j job, in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	j(in, out)
	close(out)
}

func SingleHash(in, out chan interface{}) {
	outerWg := &sync.WaitGroup{}
	for data := range in {
		fmt.Println("[SingleHash] recieved", data)
		outerWg.Add(1)
		go func(data interface{}) {
			defer outerWg.Done()
			convertedData := strconv.Itoa(data.(int))
			var result1, result2 string
			innerWg := &sync.WaitGroup{}
			innerWg.Add(2)
			go func() {
				defer innerWg.Done()
				result1 = DataSignerCrc32(convertedData)
			}()
			go func() {
				defer innerWg.Done()
				for {
					if flag := atomic.LoadUint32(&dataSignerOverheat); flag == 1 {
						time.Sleep(time.Millisecond * 100)
					} else {
						break
					}
				}
				result2 = DataSignerCrc32(DataSignerMd5(convertedData))
			}()
			innerWg.Wait()
			fmt.Println("[SingleHash] result", result1+"~"+result2)
			out <- result1 + "~" + result2
		}(data)
	}
	outerWg.Wait()
}

func MultiHash(in, out chan interface{}) {
	outerWg := &sync.WaitGroup{}

	for data := range in {
		fmt.Println("[MultiHash] recieved", data)

		outerWg.Add(1)

		go func(data interface{}) {
			defer outerWg.Done()

			result := make([]string, 6)
			convertedData := data.(string)

			mutex := &sync.Mutex{}
			innerWg := &sync.WaitGroup{}

			for th := 0; th < 6; th++ {
				innerWg.Add(1)
				go func(th int) {
					defer innerWg.Done()
					hash := DataSignerCrc32(strconv.Itoa(th) + convertedData)
					fmt.Printf("[MultiHash] %v -> %v\n", strconv.Itoa(th)+convertedData, hash)
					mutex.Lock()
					result[th] = hash
					mutex.Unlock()
				}(th)
			}

			innerWg.Wait()

			mutex.Lock()
			out <- strings.Join(result, "")
			mutex.Unlock()
		}(data)
	}

	outerWg.Wait()
}

func CombineResults(in, out chan interface{}) {
	var results []string
	for data := range in {
		fmt.Println("[Result] ->", data)
		results = append(results, data.(string))
	}
	sort.Strings(results)
	out <- strings.Join(results, "_")
}

func main() {

}
