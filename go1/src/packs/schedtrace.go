package main

import "time"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
		time.Sleep(time.Second)
	}
}
