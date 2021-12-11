package main

func main() {
	service := NewServer("127.0.0.1", 8888)
	service.Start()
}
