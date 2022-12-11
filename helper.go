package main

import (
	"fmt"
	"time"
)

func timeStamp() string {
	return fmt.Sprint(time.Now().Format("20060102150405"))
}
