package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var format = flag.String("f", "default", "time format for displaying the time")

func main() {
	flag.Parse()
	switch *format {
	case "default":
		fallthrough
	case "utc":
		fmt.Printf("%s\n", time.Now().UTC())
	case "unix":
		fmt.Printf("%d\n", time.Now().Unix())
	default:
		fmt.Print("Not a valid time fromat")
		os.Exit(1)
	}
	os.Exit(0)
}
