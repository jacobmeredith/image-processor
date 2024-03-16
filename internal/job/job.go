package job

import "image"

type Job struct {
	Input  string
	File   image.Image
	Output string
}
