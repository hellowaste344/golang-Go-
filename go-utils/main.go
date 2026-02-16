package main
import "fmt"

func main(){
	go Send(myChannel)
	fmt.Println(<-myChannel)


	fmt.Println(splitStringSlice(S))
	fmt.Println(splitAnySlice(S))

	// start goroutines
	go writeLoop(m, mu)
	go readLoop(m, mu)
	go readLoop(m, mu)
	go readLoop(m, mu)
	go readLoop(m, mu)
	go readLoop(m, mu)

	// stop program forever
	block := make(chan struct{})
	<-block
}