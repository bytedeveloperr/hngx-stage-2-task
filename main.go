package main

func main() {
	server := NewServer("0.0.0.0:3001")
	server.Run()
}
