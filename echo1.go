package main

import (
	"fmt"
	"os"
	"time"
)

func toString(i int) string {
	return fmt.Sprintf("%d", i)
}

func nanotime() int64 {
	n := time.Now()
	return n.UnixNano()
}

func main() {
	begin := nanotime()
	fmt.Println(begin)
	for key, val := range os.Args[1:] {
		fmt.Println(toString(key) + ": " + val)
	}
	end := nanotime()
	fmt.Println(begin)
	fmt.Println(end - begin)
}
