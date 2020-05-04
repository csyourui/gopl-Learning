package main

import (
	"fmt"
	"log"
	"time"

	"html"
)

//WaitForServer attempts to contact the server of a URL.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := html.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
func main() {
	fmt.Errorf("%v", WaitForServer("www.baidu.com"))
}
