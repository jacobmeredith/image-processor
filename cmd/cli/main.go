package main

import (
	"flag"
	"log"

	"github.com/jacobmeredith/image-processor/internal/job"
)

func main() {
	options := getOptions()

	jobs, err := job.Load(*options)
	if err != nil {
		log.Fatal(err)
	}

	scaled := job.Scale(jobs, *options)
	completed := job.Write(scaled, *options)

	for job := range completed {
		for _, error := range job.Errors {
			log.Printf("[%s]: %s\n", error.Step, error.Message)
		}
	}
}

func getOptions() *job.Options {
	directory := flag.String("directory", "", "the directory to read images from")
	output := flag.String("output", "", "the directory to output the images")
	verbose := flag.Bool("verbose", false, "verbose logging")
	scale := flag.Float64("scale", 0.0, "how much you want to scale the image by as a decimal e.g 0.5")

	flag.Parse()

	return &job.Options{
		InputDirectory:  *directory,
		OutputDirectory: *output,
		Verbose:         *verbose,
		Scale:           *scale,
	}
}
