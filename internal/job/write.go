package job

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func writeImage(path string, image image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, image, nil)
	if err != nil {
		return err
	}

	return nil
}

func Write(jobs <-chan Job, options Options) <-chan Job {
	new_jobs := make(chan Job)

	go func() {
		for job := range jobs {
			if options.Verbose {
				log.Printf("[Write]: %s\n", job.Input)
			}

			err := writeImage(job.Output, job.File)
			if err != nil {
				job.Errors = append(job.Errors, Error{"Write", err.Error()})
			}

			new_jobs <- job
		}

		close(new_jobs)
	}()

	return new_jobs
}
