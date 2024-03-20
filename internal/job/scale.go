package job

import (
	"image"
	"log"

	"golang.org/x/image/draw"
)

func Scale(jobs <-chan Job, options Options) <-chan Job {
	new_jobs := make(chan Job)

	if options.Scale <= 0.0 {
		return jobs
	}

	go func() {
		for job := range jobs {
			if job.Cancel {
				continue
			}

			if options.Verbose {
				log.Printf("[Scale]: %s\n", job.Input)
			}

			new_image := image.NewRGBA(image.Rect(0, 0, int(float64(job.File.Bounds().Max.X)*options.Scale), int(float64(job.File.Bounds().Max.Y)*options.Scale)))
			draw.NearestNeighbor.Scale(new_image, new_image.Rect, job.File, job.File.Bounds(), draw.Over, nil)

			job.File = new_image

			new_jobs <- job
		}

		close(new_jobs)
	}()

	return new_jobs
}
