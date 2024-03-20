package job

import (
	"image"
	"log"
	"os"
	"path/filepath"
)

func read(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func Load(options Options) (<-chan Job, error) {
	files, err := os.ReadDir(options.InputDirectory)
	if err != nil {
		return nil, err
	}

	directoryPath, err := filepath.Abs(options.InputDirectory)
	if err != nil {
		return nil, err
	}

	outputPath, err := filepath.Abs(options.OutputDirectory)
	if err != nil {
		return nil, err
	}

	jobs := make(chan Job)

	go func() {
		for _, file := range files {
			fileDirectoryPath := filepath.Join(directoryPath, file.Name())
			fileOutputPath := filepath.Join(outputPath, file.Name())

			if file.IsDir() {
				continue
			}

			if options.Verbose {
				log.Printf("[Load]: %s\n", fileDirectoryPath)
			}

			job := Job{
				Input:  fileDirectoryPath,
				Output: fileOutputPath,
				Errors: []Error{},
			}

			image, err := read(fileDirectoryPath)
			if err != nil {
				job.Errors = append(job.Errors, Error{"Load", err.Error()})
				job.Cancel = true
			} else {
				job.File = image
			}

			jobs <- job
		}

		close(jobs)
	}()

	return jobs, nil
}
