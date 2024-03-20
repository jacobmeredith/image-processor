package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jacobmeredith/image-processor/internal/job"
)

func main() {
	directory := flag.String("directory", "", "the directory to read images from")
	output := flag.String("output", "", "the directory to output the images")
	verbose := flag.Bool("verbose", false, "verbose logging")
	scale := flag.Float64("scale", 0.0, "how much you want to scale the image by as a decimal e.g 0.5")

	flag.Parse()

	options := job.Options{
		Scale: *scale,
	}

	fmt.Printf("%v", options)

	jobs, err := job.Load(*directory, *output, *verbose)
	if err != nil {
		log.Fatal(err)
	}

	scaled := job.Scale(jobs, options)
	writes := job.Write(scaled)

	for success := range writes {
		if success {
			fmt.Println("Success")
		} else {
			fmt.Println("Failed")
		}
	}
}
