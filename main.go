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
var prefix string
var suffix string
var scan string
var template string
var base string

func init() {
	current_year := time.Now().Year()
	flag.IntVar(&year, "year", current_year, "Generate folders for this year")
	flag.StringVar(&dow, "dow", "", "Comma-delimited day of weeks (Mon,Tue,Wed,Thu,Fri,Sat,Sun)")
	flag.StringVar(&subfolder, "subfolder", "", "Create subfolder")
	flag.StringVar(&format, "format", "YYYY-MM-DD", "Date format, default YYYY-MM-DD")
	flag.StringVar(&prefix, "prefix", "", "Prefix add to folder name")
	flag.StringVar(&suffix, "suffix", "", "Suffix add to folder name")
	flag.StringVar(&scan, "scan", "", "Base folder to scan folder tree")
	flag.StringVar(&template, "template", "", "Create folder tree based on template")
	flag.StringVar(&base, "base", "", "Base folder")
}

func main() {
	flag.Parse()
	checkUsage()
}

func checkUsage() {
	fmt.Printf("year = %d\n", year)
	for _, v := range strings.Split(dow, ",") {
		fmt.Printf("dow = %s\n", v)
	}
	fmt.Printf("subfolder = %q\n", subfolder)
	fmt.Printf("format = %q\n", format)
	fmt.Printf("prefix = %q\n", prefix)
	fmt.Printf("suffix = %q\n", suffix)
	fmt.Printf("scan = %q\n", scan)
	fmt.Printf("template = %q\n", template)
	fmt.Printf("base = %q\n", base)
}
