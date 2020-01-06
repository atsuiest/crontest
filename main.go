package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println("Current time in String: ", current_time.Hour(), ":", current_time.Minute(), ":", current_time.Second())
	fmt.Println(current_time.Clock())
}
