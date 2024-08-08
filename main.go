package main

func main() {
	c := make(chan int, 1)
	go StartClient(c)
	StartServer(c)
}
