package main
import "fmt"

// max can hold 100 value
var myChannel = make(chan int, 100)

func Send(ch chan int){
	ch <- 99
	close(ch)
	_, ok := <-ch
	if !ok {
		fmt.Println("Successfully closed the channel")
	}
}
