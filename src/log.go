package main

import (
	"fmt"
	"time"
)

func log(msg string) {
	fmt.Println("[" + time.Now().Format(time.RFC3339) + "] " + msg)
}
