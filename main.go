// Tom Payne calls this the tklauser coin flip.
package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	for i := 0; i < 10; i++ {
		select {
		case <-ch:
			fmt.Println("heads")
		case <-ch:
			fmt.Println("tails")
		}
	}
}
