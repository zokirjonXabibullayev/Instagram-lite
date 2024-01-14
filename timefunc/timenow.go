package timefunc

import (
	"fmt"
	"time"
)

func Timefunc() {
	var times string

	times = time.Now().Format(time.RFC850)

	fmt.Println(times)
}