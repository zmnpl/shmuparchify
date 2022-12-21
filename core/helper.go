package core

import (
	"fmt"
	"time"
)

func timeStamp() string {
	return fmt.Sprint(time.Now().Format("2006-01-02_15.04.05"))
}
