package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	fmt.Printf("%s\n", "Starting program")
	go NewFileWatcher()
	<-done
}
