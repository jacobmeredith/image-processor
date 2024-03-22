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
	cropped := job.Crop(jobs, *options)
	scaled := job.Scale(cropped, *options)
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
	cropX := flag.Int("crop-x", 0, "")
	cropY := flag.Int("crop-y", 0, "")
	cropWidth := flag.Int("crop-width", 0, "")
	cropHeight := flag.Int("crop-height", 0, "")
	quality := flag.Int("quality", 100, "")

	flag.Parse()

	return &job.Options{
		InputDirectory:  *directory,
		OutputDirectory: *output,
		Verbose:         *verbose,
		Scale:           *scale,
		Quality:         *quality,
		Crop: job.CropOptions{
			X:      *cropX,
			Y:      *cropY,
			Width:  *cropWidth,
			Height: *cropHeight,
		},
	}
}
