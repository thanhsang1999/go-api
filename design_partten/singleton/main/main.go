package main

import (
	"go-api/design_partten/singleton"
	"log"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			log.Printf("%p\n", singleton.GetInstance())
		}()
	}
	time.Sleep(10 * time.Second)
}
