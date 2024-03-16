package job

import (
	"image"

	"golang.org/x/image/draw"
)

func Resize(jobs <-chan Job) <-chan Job {
	new_jobs := make(chan Job)

	go func() {
		for job := range jobs {
			new_image := image.NewRGBA(image.Rect(0, 0, job.File.Bounds().Max.X/2, job.File.Bounds().Max.Y/2))
			draw.NearestNeighbor.Scale(new_image, new_image.Rect, job.File, job.File.Bounds(), draw.Over, nil)

			job.File = new_image

			new_jobs <- job
		}

		close(new_jobs)
	}()

	return new_jobs
}
