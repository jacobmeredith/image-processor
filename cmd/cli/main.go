package main

import (
	"flag"
	"log"

	"github.com/jacobmeredith/image-processor/internal/job"
)

func main() {
	options := getFlags()

	jobs, err := job.Load(options.InputDirectory, options.OutputDirectory, options.Verbose)
	if err != nil {
		log.Fatal(err)
	}

	scaled := job.Scale(jobs, *options)
	job.Write(scaled)
}

func getFlags() *job.Options {
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
