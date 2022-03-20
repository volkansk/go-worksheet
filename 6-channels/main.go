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

	// fmt.Println(<-c) //wait for "a" value to be sent into this channel
	// fmt.Println(<-c) //wait for second go routine result

	//received all go routine result
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(l, c) //infinitive loop
	// }

	//function literal
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // -> blocking code!

	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down i think" //send the value into this channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "Yep its up"
	c <- link
}
