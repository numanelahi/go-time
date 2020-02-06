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
	case "ymd":
		year, month, day := time.Now().Date()
		fmt.Printf("%d-%d-%d\n", year, month, day)
	case "year", "y":
		fmt.Printf("%d\n", time.Now().Year())
	case "month", "m":
		fmt.Printf("%s\n", time.Now().Month().String())
	case "day", "d":
		fmt.Printf("%d\n", time.Now().Day())
	case "ym":
		fmt.Printf("%d-%d\n", time.Now().Year(), time.Now().Month())
	case "utc":
		fmt.Printf("%s\n", time.Now().UTC())
	case "unix":
		fmt.Printf("%d\n", time.Now().Unix())
	default:
		fmt.Println("Not a valid time fromat")
		os.Exit(1)
	}
	os.Exit(0)
}
