package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	
}

func checkLink(link string,c chan string){
	_, err := http.Get(link)

	if err !=nil {
		fmt.Println(link , "might be down !")
		c <- "Might be down I Think"
		return 
	}
	fmt.Println(link,"is up !")
	c <- "Yep it is up ! "

}