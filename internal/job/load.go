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

func Load(directory, output string, verbose bool) (<-chan Job, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	directoryPath, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}

	outputPath, err := filepath.Abs(output)
	if err != nil {
		return nil, err
	}

	jobs := make(chan Job)

	go func() {
		for _, file := range files {
			fileDirectoryPath := filepath.Join(directoryPath, file.Name())
			fileOutputPath := filepath.Join(outputPath, file.Name())

			if verbose {
				log.Printf("Reading: %s\n", fileDirectoryPath)
			}

			if file.IsDir() {
				continue
			}

			image, err := read(fileDirectoryPath)
			if err != nil {
				continue
			}

			job := Job{
				Input:  fileDirectoryPath,
				File:   image,
				Output: fileOutputPath,
			}

			jobs <- job
		}

		close(jobs)
	}()

	return jobs, nil
}
