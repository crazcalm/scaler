package main

import (
	"fmt"
	"flag"
)

var height = flag.Float64("h", 0, "Image height")
var width = flag.Float64("w", 0, "Image width" )
var scaler = flag.Float64("s", 1.0, "The scaler used on the passed in values")
var verbose = flag.Bool("v", false, "Verbose output")


func main (){
	flag.Parse()
	if *verbose {
		fmt.Printf("Width: %.0f, Height: %.0f, Scaler: %f\n", *width, *height, *scaler)
	}
	fmt.Printf("%.0f %.0f\n", *width * *scaler , *height * *scaler)
}