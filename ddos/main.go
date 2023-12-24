package main

import (
	"dict/ddos/ddos"
	"fmt"
	"time"
)

func main() {
	fmt.Println("DDOS..........")
	Ddos()

}
func Ddos() {
	workers := 10000000
	d, err := ddos.New("http://localhost:8080/get-message/Apple", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: http://localhost:8080/get-message/Apple")
}
