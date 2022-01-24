package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){

	links := []string{
		"http://ramzinex.com",
		"http://digikala.com",
		"http://nobitex.ir",
		"http://snapp.ir",
	}

	c:=make(chan string)
	

	for _,link:= range links{
		go checkLink(link, c)
	}
// Adding time.sleep to this for loop will stop main routin
// So we create function literal and make timer on there
	for l:=range c {
		go func(link  string){
		time.Sleep(5 * time.Second)
		checkLink(link,c)
		}(l)
	}

}

func checkLink(link string,c chan string){
	_, err := http.Get(link)

	if err !=nil {
		fmt.Println(link , "might be down !")
		c <- link
		return 
	}
	fmt.Println(link,"is up !")
	c <- link

}