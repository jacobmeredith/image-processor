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

	flag.Parse()

	jobs, err := job.Load(*directory, *output, *verbose)
	if err != nil {
		log.Fatal(err)
	}

	resized := job.Resize(jobs)
	writes := job.Write(resized)

	for success := range writes {
		if success {
			fmt.Printf("Success")
		} else {
			fmt.Println("Failed")
		}
	}
}
