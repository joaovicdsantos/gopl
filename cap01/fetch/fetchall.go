package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	file, err := os.Create("fetchall.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, url := range os.Args[1:] {
		go fetch(url, ch, file)
	}

	for range os.Args[1:] {
		res := <-ch
		file.Write([]byte(res))
		fmt.Println(res)
		file.Sync()
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, file *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	response, _ := io.ReadAll(resp.Body)

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	file.WriteString(fmt.Sprintf("\n-----------------------------\n%s\n-----------------------------\n\n", url))

	if err != nil {
		res := fmt.Sprintf("while reading %s: %v\n", url, err)
		file.WriteString(res)
		ch <- res
		return
	}
	secs := time.Since(start).Seconds()

	file.Write(response)
	res := fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	file.WriteString(res)
	ch <- res

}
