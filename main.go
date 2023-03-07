package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool); // make channel which has bool values
	people := [2]string{"hyuk", "sun"};

	for _, person := range people {
		go isSexy(person, channel);
	}

	fmt.Println("Recieved this message: ", <- channel);
	fmt.Println("Recieved this message: ", <- channel);
}

func isSexy(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	channel <- true; // send value to channel.
}