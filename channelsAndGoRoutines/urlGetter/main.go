package main

import (
	"fmt"
	"net/http"
	"time"
)

func getLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Seems like link %s is down\n", link)
		//c <- fmt.Sprintf("Seems like link %s is down\n", link)
		c <- link
		return
	}
	fmt.Printf("%s alive\n", link)
	//c <- fmt.Sprintf("link %s is alive\n", link)
	c <- link
}

func main() {
	links := []string{
		"http://www.docker.com",
		"http://www.google.com",
		"https://aws.amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, val := range links {
		go getLink(val, c)
	}

	for l := range c {
		// never try to access the values of the parent routine
		// instead, copy the value and send it in the child routine
		go func(link string) {
			time.Sleep(2 * time.Second)
			getLink(link, c)
		}(l)
	}
}

// import "fmt"

// func main() {
// 	c := make(chan string)
// 	c <- "Dcc@MOOMOO"
// 	v, _ := <-c
// 	close(c)
// 	fmt.Println(v)
// }
