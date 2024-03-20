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

func Write(jobs <-chan Job) <-chan bool {
	status := make(chan bool)

	go func() {
		for job := range jobs {
			log.Printf("Writing: %s\n", job.Input)
			err := writeImage(job.Output, job.File)
			if err != nil {
				status <- false
				continue
			}

			status <- true
		}

		close(status)
	}()

	return status
}
