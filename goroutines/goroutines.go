package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const filename string = "links.txt"

func getContent(urlChan chan string) {
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error with opening the file", err)
		return
	}
	defer fileObj.Close()

	contents := bufio.NewScanner(fileObj)
	for contents.Scan() {
		urlChan <- contents.Text()
	}
	close(urlChan)
}

func readStatus(ch chan string, wg *sync.WaitGroup) {
	for URL := range ch {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(URL)
			if err != nil {
				fmt.Println("Error with opening the file", err)
				return
			}
			res.Body.Close()

			statusCode := res.StatusCode
			if statusCode == 200 {
				fmt.Println(statusCode, ":The site is up and running", url)
			} else {
				fmt.Println(statusCode, "The site needs to checked")
			}
		}(URL)
	}
}

func main() {
	startNow := time.Now()

	ch := make(chan string)
	var wg sync.WaitGroup

	go getContent(ch)
	readStatus(ch, &wg)

	wg.Wait()

	fmt.Println("Time taken:", time.Since(startNow))
}
