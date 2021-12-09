package config

import "testing"

import "fmt"

func ping(pings chan<- string, msg string) {  // 只写方向
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {  // 只读， 只写
	msg := <-pings
	pongs <- msg
}

func TestRunChannel(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
