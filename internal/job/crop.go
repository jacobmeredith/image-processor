package job

import (
	"image"
	"log"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Crop(jobs <-chan Job, options Options) <-chan Job {
	new_jobs := make(chan Job)

	if options.Crop.Width == 0 || options.Crop.Height == 0 {
		return jobs
	}

	go func() {
		for job := range jobs {
			if job.Cancel {
				continue
			}

			if options.Verbose {
				log.Printf("[Crop]: %s\n", job.Input)
			}

			cropSize := image.Rect(0, 0, options.Crop.Width, options.Crop.Height).Add(image.Point{options.Crop.X, options.Crop.Y})

			job.File = job.File.(SubImager).SubImage(cropSize)

			new_jobs <- job
		}

		close(new_jobs)
	}()

	return new_jobs
}
