package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//   Also blocking call, lec 75
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	//for i := 0; i < len(links); i++ {
	//fmt.Println(<-c)
	//}

	//for {
	// go's logic realizes that  <-c will eventually return a string
	//go checkLink(<-c, c)
	//}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	//            Blocking call!
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		//c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "Yep, it's up"
	c <- link
}
