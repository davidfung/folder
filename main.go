package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var year int
var dow string
var subfolder string
var format string

func init() {
	current_year := time.Now().Year()
	flag.IntVar(&year, "year", current_year, "Generate folders for this year")
	flag.StringVar(&dow, "dow", "", "Comma-delimited day of weeks (Mon,Tue,Wed,Thu,Fri,Sat,Sun)")
	flag.StringVar(&subfolder, "subfolder", "", "Create subfolder")
	flag.StringVar(&format, "format", "YYYY-MM-DD", "Date format, default YYYY-MM-DD")
}

func main() {
	flag.Parse()
	fmt.Printf("year = %d\n", year)
	for _, v := range strings.Split(dow, ",") {
		fmt.Printf("dow = %s\n", v)
	}
	fmt.Printf("subfolder = %s\n", subfolder)
	fmt.Printf("format = %s\n", format)
}
