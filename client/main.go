// concurrent.go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MakeRequest(baseURL string, streamId int, ch chan<- string) {

	start := time.Now()

	n := 0
	for n < 10000 {
		n++
		resp, _ := http.Get(baseURL)
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		fmt.Println("Stream ", streamId, "iteration ", n)
		fmt.Printf("%s", body)
		//time.Sleep(1 * time.Millisecond)
		//bodyLen += len(body)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f elapsed with url %s", secs, baseURL)

}

func main() {

	urlPtr := flag.String("url", "http://www.sohu.com", "base url")
	streamsPtr := flag.Int("streams", 1, "number of streams")
	flag.Parse()
	baseURL := *urlPtr
	streams := *streamsPtr

	start := time.Now()
	ch := make(chan string)

	n := 0
	for n < streams {
		go MakeRequest(baseURL, n, ch)
		n++
	}

	fmt.Println(<-ch)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
