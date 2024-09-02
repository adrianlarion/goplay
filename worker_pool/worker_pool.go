package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//urls to process
	urls := make([]string, 7)

	//number of workers
	bigWPoolNum := 5

	//jobs channel
	chUrl := make(chan string)

	//results channel
	resCh := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(bigWPoolNum)

	for i := 0; i < bigWPoolNum; i++ {
		go func() {
			defer wg.Done()
			//get urls from jobs channel
			for v := range chUrl {
				//do work
				time.Sleep(1 * time.Second)
				fmt.Println("looking for url ", v)
				resCh <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	go func() {
		//push urls to jobs channel
		for _, v := range urls {
			chUrl <- v
		}
		fmt.Println("closing churl")
		close(chUrl)

	}()

	//get results from results channel
	for res := range resCh {
		fmt.Println("res is ", res)
	}
}
