package job

import "image"

type Options struct {
	Scale float64
}

type Job struct {
	Input  string
	File   image.Image
	Output string
}
