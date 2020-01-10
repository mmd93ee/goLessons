package main

import (
  "net/http"
  "fmt"
  "time"
  "os"
  "io"
  "io/ioutil"
)



func main() {
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] {
      go fetch(url, ch) // start a go routine
    }
    for range os.Args[1:] {
      fmt.Println(<-ch) // recieve from channel 'ch'
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string){
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) // Send to channel
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()
  if err != nil {
    ch <- fmt.Sprintf("Error wile reading: %s : %s", url, err)
    return
  }

  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
