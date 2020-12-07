package main

func main() {
	c := New("inmemory")
	NewServer(c).Listen()
}
