package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string); // make channel which has bool values
	people := [2]string{"hyuk", "sun"};

	for _, person := range people {
		go isSexy(person, channel);
	}
	
	fmt.Println("Waiting for messages");
	
	for i := 0; i < len(people); i++ {
		fmt.Println("Recieved this message: ", <- channel);
	}

	fmt.Println("end!");
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"; // send value to channel.
}